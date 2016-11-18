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

// OK 0
const (
	OK int = iota
	ERROR
)

type res struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

func wikiHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<html><head><title>Remind</title></head><body>")
	fmt.Fprintln(w, "<h1>Remind</h1>")
	fmt.Fprintln(w, "<ul><li>GET</li><li>POST</li><li>PUT</li><li>DELETE</li></ul>")
	fmt.Fprintln(w, "</body></html>")
}

func indexHandle(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w, req.RequestURI)
	// io.WriteString(w, req.URL.Path)

	msg := "[" + req.Method + "]" + req.URL.Path
	ss := strings.Split(req.URL.Path, "/")[1:]
	ns := len(ss)
	data := make(map[string]interface{})
	p := 0
	if ns > p {
		data["app"] = ss[0]
		p++
	}
	if ns > p {
		data["controller"] = ss[1]
		p++
	}
	if ns > p {
		data["action"] = ss[2]
		p++
	}
	if ns > p && ns%2 == 0 {
		data["version"] = ss[3]
		p++
	}
	for ; p < ns; p += 2 {
		data[ss[p]] = ss[p+1]
	}
	result := res{ERROR, msg, data}
	//u, err := url.Parse(r)
	switch req.Method {
	case "GET":
		result.Code = OK
	case "POST":
		contentType := req.Header.Get("Content-Type")
		result.Msg += "--" + contentType
		body, _ := ioutil.ReadAll(req.Body)
		switch contentType {
		case "application/json":
			var mapBody map[string]interface{}
			json.Unmarshal(body, &mapBody)
			// io.WriteString(w, mapBody["user"].(string))
			result.Data["JSON"] = mapBody
		case "application/x-www-form-urlencoded;charset=utf-8":
			err := req.ParseForm()
			if err != nil {
				fmt.Println(err)
			}
			// io.WriteString(w, req.Form["user"][0])
			// io.WriteString(w, req.FormValue("plan"))
			m := make(map[string]string)
			fmt.Println(req.Form)
			for k, v := range req.Form {
				// io.WriteString(w, k+v[0])
				fmt.Printf("%s : %s\n", k, v[0])
				m[k] = v[0]
			}
			result.Data["FORM"] = req.Form
		default:
			result.Data["BODY"] = string(body)

		}

	}
	j, _ := json.Marshal(result)
	fmt.Println(result)
	io.WriteString(w, string(j))
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
