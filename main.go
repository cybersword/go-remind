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

	_ "github.com/go-sql-driver/mysql"
)

type res struct {
	Code int         `json:"code"`
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
	utils.Notice(msg)
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
	} else {
		params["version"] = "master"
	}
	for ; p < ns; p += 2 {
		params[ss[p]] = ss[p+1]
	}
	result := res{utils.ERROR, msg, nil}
	req.ParseForm() // parse params in POST|PUT|PATCH body form and params in query
	form := make(map[string]string)
	for fk, fv := range req.Form {
		form[fk] = fv[0]
	}
	params["FORM"] = form
	params["METHOD"] = req.Method
	switch req.Method {
	case "GET":
		result.Code = utils.OK
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

	vApp, ok1 := params["app"]
	vc, ok2 := params["controller"]
	if !ok1 {
		result.Msg = "app not set"
	} else if !ok2 {
		result.Msg = "controller not set"
	} else {
		// init app
		appName := strings.ToLower(vApp.(string))
		c := strings.Title(vc.(string))
		c = strings.Replace(c, " ", "", -1)
		c = strings.Replace(c, "-", "", -1)
		c = strings.Replace(c, "_", "", -1)
		appInstance := &app.App{Name: appName}
		v := reflect.ValueOf(appInstance)
		t := reflect.TypeOf(appInstance)
		// call controller
		// TODO: 要设计一下不同的响应方式
		_, ok := t.MethodByName(c)
		if !ok {
			result.Msg = "controller not found"
		} else {
			// 暂且特殊处理
			if c == "Hi" {
				message := appInstance.Hi(params)
				utils.Notice(message)
				io.WriteString(w, message)
				return

			}
			f := v.MethodByName(c)
			in := []reflect.Value{reflect.ValueOf(params)}
			ret := f.Call(in)
			result.Code = int(ret[0].Int()) // int64 -> int
			result.Msg = ret[1].String()
			result.Data = ret[2].Interface()
		}

	}

	// response
	j, _ := json.Marshal(result)
	utils.Notice(result)
	io.WriteString(w, string(j))
}

func main() {

	utils.Notice("启动检测-Notice")
	utils.Fatal("启动检测-Fatal")

	utils.DangDang("http://product.dangdang.com/23910258.html")
	utils.DangDang("http://product.dangdang.com/23800641.html")

	// 查看接口文档
	http.HandleFunc("/wiki", wikiHandle)
	// 查看当前配置
	http.HandleFunc("/conf", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Conf:\n"+req.URL.Path[1:]+"\n")
	})
	// 健康检查
	http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Health Check:\n"+req.URL.Path[1:]+"\n")
	})
	http.HandleFunc("/", indexHandle)
	err := http.ListenAndServe(":8034", nil) // always returns a non-nil error.
	log.Fatal("ListenAndServe: ", err.Error())
}
