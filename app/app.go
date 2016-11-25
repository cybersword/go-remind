package app

import (
	"log"
)

type App struct {
	Name string
}

func (a *App) Plan(params map[string]interface{}) (int, string, interface{}) {
	log.Println("call Plan")
	return 0, a.Name + "call plan", params
}

func (a *App) Task(params map[string]interface{}) (int, string, interface{}) {
	log.Println("call Task")
	return 0, a.Name + "call task", params
}
