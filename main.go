package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func wikiHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<html><head><title>Remind</title></head><body>")
	io.WriteString(w, "<h1>Remind ...</h1>\n")
	fmt.Fprintln(w, "</body></html>")
}

func indexHandle(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, req.RequestURI)
	io.WriteString(w, req.URL.Path)
	ss := strings.Split(req.URL.Path, "/")
	for _, s := range ss {
		io.WriteString(w, s+"\n")
	}
	//u, err := url.Parse(r)
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
			// io.WriteString(w, req.Form["user"][0])
			// io.WriteString(w, req.FormValue("plan"))
			for k, v := range req.Form {
				// io.WriteString(w, k+v[0])
				fmt.Fprintf(w, "%s : %s\n", k, v[0])
			}
		}

	}
}

func main() {
	http.HandleFunc("/wiki", wikiHandle)
	// 最长匹配原则
	http.HandleFunc("/bar", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, "+req.URL.Path[1:]+"\n")
	})
	http.HandleFunc("/", indexHandle)
	err := http.ListenAndServe("localhost:8765", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
