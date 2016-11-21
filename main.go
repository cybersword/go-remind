package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"strings"

	"github.com/cybersword/go-remind/utils"
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
	// init params
	msg := "[" + req.Method + "]" + req.URL.Path
	ss := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	ns := len(ss)
	params := make(map[string]interface{})
	p := 0
	arrRoute := []string{"app", "controller", "action"}
	for ; p < ns; p++ {
		params[arrRoute[p]] = ss[p]
	}
	if p < ns && ns%2 == 0 {
		params["version"] = ss[p]
		p++
	}
	for ; p < ns; p += 2 {
		params[ss[p]] = ss[p+1]
	}
	result := res{ERROR, msg, nil}
	req.ParseForm() // parse params in POST|PUT|PATCH body form and params in query
	fmt.Println(req.Form)
	params["FORM"] = req.Form
	params["METHOD"] = req.Method
	switch req.Method {
	case "GET":
		result.Code = OK
	case "PUT":
		fallthrough
	case "PATCH":
		fallthrough
	case "POST":
		ct := req.Header.Get("Content-Type")
		result.Msg += " -- " + ct
		body, _ := ioutil.ReadAll(req.Body)
		ct, _, _ = mime.ParseMediaType(ct)
		switch ct {
		case "application/json":
			var m map[string]interface{}
			json.Unmarshal(body, &m)
			params["JSON"] = m
		case "application/x-www-form-urlencoded":
			// 已经通过 ParseForm 解析过了
		default:
			// 其它不支持的类型
			params["BODY"] = string(body)
		}
	}

	// call process func here
	fmt.Println(params)
	j, _ := json.Marshal(result)
	fmt.Println(result)
	io.WriteString(w, string(j))
}

func main() {
	logFileName := "debug.log"
	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Fatalln("open log file error !")
	}
	defer logFile.Close()

	sl := utils.GetSimpleLogger()
	sl.Notice("我是notice1")
	sl.Fatal("bugbugbug")
	sl.Notice("我是notice2")
	debugLog := log.New(logFile, "[Debug]", log.Lshortfile|log.LstdFlags)
	debugLog.Println("A debug message here")
	debugLog.SetPrefix("[Info]")
	debugLog.Println("A Info Message here ")
	debugLog.SetFlags(debugLog.Flags() | log.Lmicroseconds)
	debugLog.Println("A different prefix")
	http.HandleFunc("/wiki", wikiHandle)
	// 最长匹配原则
	http.HandleFunc("/bar", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, "+req.URL.Path[1:]+"\n")
	})
	http.HandleFunc("/", indexHandle)
	err = http.ListenAndServe("localhost:8765", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
