package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// helloHandle just say hello
func helloHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintln(w, "<html><body>")
	fmt.Fprintf(w, "Hello, "+req.URL.Path[1:]+"\n")
	io.WriteString(w, "<h1>hello, world</h1>\n")
	fmt.Fprintln(w, "</body></html>")
}

func indexHandle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		io.WriteString(w, `{"code": 0, "msg": "GET"}`)
	case "POST":
		contentType := req.Header.Get("Content-Type")
		io.WriteString(w, contentType+"\n")

		switch contentType {
		case "application/json":
			var mapBody map[string]interface{}
			body, _ := ioutil.ReadAll(req.Body)
			json.Unmarshal(body, &mapBody)
			io.WriteString(w, mapBody["user"].(string))
		default:
			req.ParseForm()
			io.WriteString(w, req.Form["user"][0])
			io.WriteString(w, req.FormValue("plan"))
		}

	}
}

func main() {
	http.HandleFunc("/hello", helloHandle)
	// 最长匹配原则
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/", indexHandle)
	err := http.ListenAndServe("localhost:8765", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
