package main

import "fmt"
import "reflect"
import "encoding/xml"

import "./app"
import "log"
import "strings"

type st struct {
}

func (this *st) Echo() {
	fmt.Println("echo()")
}

func (this *st) Echo2() {
	fmt.Println("echo--------------------()")
}

var xmlstr string = `<root>  
<func>Echo</func>  
<func>Echo2</func>  
</root>`

type st2 struct {
	E []string `xml:"func"`
}

func main() {
	s2 := st2{}
	xml.Unmarshal([]byte(xmlstr), &s2)

	s := &st{}
	v := reflect.ValueOf(s)

	v.MethodByName(s2.E[1]).Call(nil)
	dawn := reflect.ValueOf(&app.App{"dawn"})
	f := dawn.MethodByName(strings.Title("plan"))
	log.Println(f)
	//     dawn := reflect.New(&app.App{"dawn"})
	params := map[string]interface{}{"a": 2, "b": "cccccc"}
	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(params)
	log.Println(params)
	log.Println("ccc")
	log.Println(dawn)
	log.Println(in)
	f.Call(in)

}
