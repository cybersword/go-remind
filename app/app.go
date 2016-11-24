package app

import (
	"log"
)

type App struct {
	Name string
}

func (a *App) Plan(params map[string]interface{}) {
	log.Println("call Plan")
}

func (a *App) Task(params map[string]interface{}) {
	log.Println("call Task")
}
