package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
)

const keyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("Got into getRoot with the server address ==> %s\n", ctx.Value(keyServerAddr))
	fmt.Fprint(w, "This is the root of my website")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	page := r.URL.Query().Get("page")

	w.WriteHeader(http.StatusOK)
	fmt.Printf("Got into getHello with the server address ==> %s\n", ctx.Value(keyServerAddr))
	fmt.Fprintf(w, "Hello world with the page as %s", page)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	ctx, cancelCtx := context.WithCancel(context.Background())

	go func() {
		server := &http.Server{
			Addr:    ":3333",
			Handler: mux,
			BaseContext: func(l net.Listener) context.Context {
				ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())

				return ctx
			},
		}

		err := server.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("The server is closed")
		} else if err != nil {
			fmt.Printf("A server error occured %s\n", err)
		}

		cancelCtx()
	}()

	// We need this so that we don't exist the main program immediately before the go routine finishes
	<-ctx.Done()
}
