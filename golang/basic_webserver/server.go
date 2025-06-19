package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func incrementCounter(w http.ResponseWriter, _r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, "The counter is %s", strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// 	// http.ServeFile(w, r, r.URL.Path[1:])
	// })

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	fmt.Println("Listening to requests")
	// log.Fatal(http.ListenAndServe(":8081", nil))
	log.Fatal(http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil))
}
