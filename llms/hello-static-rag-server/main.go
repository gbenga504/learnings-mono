package main

import (
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores/weaviate"
)

const modelName = "llama3.1:8b"

func main() {
	ctx := context.Background()

	// Create a new ollama client with the specified model.
	ollamaClient, err := ollama.New(ollama.WithModel((modelName)))

	if err != nil {
		log.Fatal(err)
	}

	// Create a new embedder with the specified ollama client.
	emb, err := embeddings.NewEmbedder(ollamaClient)

	if err != nil {
		log.Fatal(err)
	}

	// Create a new Weaviate client with the specified embedder.
	wvStore, err := weaviate.New(weaviate.WithEmbedder(emb), weaviate.WithScheme("http"), weaviate.WithHost("localhost:"+cmp.Or(os.Getenv("WEAVIATE_PORT"), "8080")), weaviate.WithIndexName("Document"))

	if err != nil {
		log.Fatal(err)
	}

	server := &ragServer{
		ctx:          ctx,
		wvStore:      wvStore,
		ollamaClient: ollamaClient,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /add/", server.addDocumentsHandler)
	mux.HandleFunc("POST /query/", server.queryHandler)

	port := cmp.Or(os.Getenv("PORT"), "6000")
	address := "localhost:" + port

	log.Println("Listening on", address)
	log.Fatal(http.ListenAndServe(address, mux))
}

type ragServer struct {
	ctx          context.Context
	wvStore      weaviate.Store
	ollamaClient *ollama.LLM
}

func (rs *ragServer) addDocumentsHandler(w http.ResponseWriter, req *http.Request) {
	// Parse HTTP request from JSON
	type document struct {
		Text string `json:"text"`
	}

	type addRequest struct {
		Documents []document `json:"documents"`
	}

	ar := &addRequest{}

	err := readRequestJSON(req, ar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// store documents and their embeddings in weaviate
	var wvDocs []schema.Document
	for _, doc := range ar.Documents {
		wvDocs = append(wvDocs, schema.Document{PageContent: doc.Text})
	}

	_, err = rs.wvStore.AddDocuments(rs.ctx, wvDocs)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (rs *ragServer) queryHandler(w http.ResponseWriter, req *http.Request) {
	// Parse HTTP request from JSON
	type queryRequest struct {
		Content string `json:"content"`
	}

	qr := &queryRequest{}
	err := readRequestJSON(req, qr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find the most similar documents
	docs, err := rs.wvStore.SimilaritySearch(rs.ctx, qr.Content, 3)
	if err != nil {
		http.Error(w, fmt.Errorf("similarity search: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	var docsContents []string
	for _, doc := range docs {
		docsContents = append(docsContents, doc.PageContent)
	}

	// Create a RAG query for the LLM with the most relevant documents as context
	ragQuery := fmt.Sprintf(ragTemplateStr, qr.Content, strings.Join(docsContents, "\n"))
	respText, err := llms.GenerateFromSinglePrompt(rs.ctx, rs.ollamaClient, ragQuery, llms.WithModel(modelName))

	if err != nil {
		log.Printf("calling generative model: %v", err.Error())
		http.Error(w, "generative model error", http.StatusInternalServerError)
		return
	}

	renderJSON(w, respText)
}

const ragTemplateStr = `
I will ask you a question and will provide some additional context information.
Assume this context information is factual and correct, as part of internal
documentation.
If the question relates to the context, answer it using the context.
If the question does not relate to the context, answer it as normal.

For example, let's say the context has nothing in it about tropical flowers;
then if I ask you about tropical flowers, just answer what you know about them
without referring to the context.

For example, if the context does mention minerology and I ask you about that,
provide information from the context along with general knowledge.

Question:
%s

Context:
%s
`
