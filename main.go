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

type res struct {
	code int
	msg  string
	data map[string]interface{}
}

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
		body, _ := ioutil.ReadAll(req.Body)
		switch contentType {
		case "application/json":
			var mapBody map[string]interface{}

			json.Unmarshal(body, &mapBody)
			io.WriteString(w, mapBody["user"].(string))
		case "application/x-www-form-urlencoded;charset=utf-8":

			req.ParseForm()
			// io.WriteString(w, req.Form["user"][0])
			// io.WriteString(w, req.FormValue("plan"))
			var m map[string]interface{}
			for k, v := range req.Form {
				// io.WriteString(w, k+v[0])
				fmt.Fprintf(w, "%s : %s\n", k, v[0])
				m[k] = v[0]
			}
			r := res{0, "form", m}
			j, _ := json.Marshal(r)
			io.WriteString(w, string(j))
		default:
			fmt.Fprintf(w, string(body))

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
