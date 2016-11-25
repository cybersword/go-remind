package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"reflect"
	"strings"

	"github.com/cybersword/go-remind/app"
	"github.com/cybersword/go-remind/utils"
)

// OK 0
const (
	OK int64 = iota
	ERROR
)

type res struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func wikiHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<html><head><title>Remind</title></head><body>")
	fmt.Fprintln(w, "<h1>Remind</h1>")
	fmt.Fprintln(w, "<ul><li>GET</li><li>POST</li><li>PUT</li><li>DELETE</li></ul>")
	fmt.Fprintln(w, "</body></html>")
}

func indexHandle(w http.ResponseWriter, req *http.Request) {
	msg := "[" + req.Method + "]" + req.URL.Path
	// io.WriteString(w, req.RequestURI)
	// io.WriteString(w, req.URL.Path)
	// init params
	params := make(map[string]interface{})
	ss := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	ns := len(ss)
	rr := []string{"app", "controller", "action"}
	nr := len(rr)
	p := 0
	for ; p < ns && p < nr; p++ {
		params[rr[p]] = ss[p]
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
	vc, ok := params["controller"]
	if !ok {
		result.Msg = "hello?"
	} else {
		c := strings.Title(vc.(string))
		dawn := &app.App{Name: "dawn"}
		v := reflect.ValueOf(dawn)
		t := reflect.TypeOf(dawn)
		_, ok := t.MethodByName(c)
		if !ok {
			result.Msg = "controller not found"
		} else {
			f := v.MethodByName(c)

			in := []reflect.Value{reflect.ValueOf(params)}
			ret := f.Call(in)
			result.Code = ret[0].Int()
			result.Msg = ret[1].String()
			result.Data = ret[2].Interface()
		}
	}

	j, _ := json.Marshal(result)
	utils.Notice(result)
	io.WriteString(w, string(j))
}

func main() {

	utils.Notice("我是notice")
	utils.Fatal("error static func")
	utils.Notice("第二条日志")

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
