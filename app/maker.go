package app

import "github.com/cybersword/go-remind/utils"

// Maker ...
type Maker struct {
	AppName string
}

// Name ...
func (app *Maker) Name() string {
	return "Dog"
}

// VideoInfo ..
func (a *Maker) VideoInfo(params map[string]interface{}) (int, string, interface{}) {
	vid, ok := params["vid"].(string)
	if !ok {
		errMsg := "缺少vid参数"
		utils.Fatal(errMsg, params)
		return utils.ERROR, errMsg, params
	}

	return utils.OK, "vid ok", vid
}
