package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/cybersword/string"
)

// HelloServer just say hello
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello, "+req.URL.Path[1:]+"\n")
	s := req.URL.Path[1:]
	fmt.Fprintf(w, "reverse: "+string.Reverse(s))
}

func main() {
	http.HandleFunc("/", HelloServer)
	// 最长匹配原则
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	err := http.ListenAndServe("localhost:8765", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
