package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// HelloServer just say hello
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/", HelloServer)
	// 最长匹配原则
	http.HandleFunc("/foo", fooHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	err := http.ListenAndServe("localhost:8765", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
