package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// //
// // Middleware 1
// //
type Logger struct {
	handler http.Handler
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

func NewLogger(handlerToWrap http.Handler) Logger {
	return Logger{handlerToWrap}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	wrappedMux := NewLogger(mux)

	log.Fatal(http.ListenAndServe(":3333", wrappedMux))
}
