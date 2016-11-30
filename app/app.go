package app

import (
	"github.com/cybersword/go-remind/utils"
)

// App entry
type App struct {
	Name string
}

// Plan controller
func (a *App) Plan(params map[string]interface{}) (int, string, interface{}) {
	utils.Notice("call Plan")
	return 0, a.Name + "call plan", params
}

// Task controller
func (a *App) Task(params map[string]interface{}) (int, string, interface{}) {
	utils.Notice("call Task")
	return 0, a.Name + "call task", params
}
