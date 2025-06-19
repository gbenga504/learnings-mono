package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var articles = []Article{
	{Id: "1", Title: "Hello", Desc: "Article description", Content: "Article content"},
	{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article content"},
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to our REST Api")
	}).Methods("GET")

	myRouter.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(articles)
	}).Methods("GET")

	myRouter.HandleFunc("/articles/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		articleId := vars["id"]

		for _, article := range articles {

			if article.Id == articleId {
				json.NewEncoder(w).Encode(article)
				break
			}
		}
	}).Methods("GET")

	myRouter.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := io.ReadAll(r.Body)
		var newArticle Article

		json.Unmarshal(reqBody, &newArticle)

		articles = append(articles, newArticle)

		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(newArticle)
	}).Methods("POST")

	myRouter.HandleFunc("/articles/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		articleId := vars["id"]

		for index, article := range articles {
			if article.Id == articleId {
				articles = append(articles[:index], articles[index+1:]...)
			}
		}

		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(articles)
	}).Methods("DELETE")

	myRouter.HandleFunc("/articles/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		articleId := vars["id"]

		reqBody, _ := io.ReadAll(r.Body)
		var updates Article

		json.Unmarshal(reqBody, &updates)

		for index, article := range articles {
			if article.Id == articleId {
				articles[index] = updates
			}
		}

		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(articles)
	}).Methods("PUT")

	myRouter.HandleFunc("/articles/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		articleId := vars["id"]

		reqBody, _ := io.ReadAll(r.Body)

		for index, article := range articles {
			if article.Id == articleId {
				var mergedArticle Article

				// Merge the updates with the already existing article
				// We json stringify the article then parse it into the merged article
				// Then we go ahead to parse the req body also
				///////////////////
				article, _ := json.Marshal(articles[index])
				json.Unmarshal(article, &mergedArticle)

				json.Unmarshal(reqBody, &mergedArticle)
				////////////////
				////////////////

				articles[index] = mergedArticle
			}
		}

		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(articles)
	}).Methods("PATCH")

	fmt.Println("Running server on 8081 port")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
