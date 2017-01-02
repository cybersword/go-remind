package app

import "github.com/cybersword/go-remind/utils"

// App entry
type App struct {
	Name string
}

// controllers

// Hi /lab/hi/alpha token is gk2017
func (a *App) Hi(params map[string]interface{}) string {

	form, ok := params["FORM"].(map[string]string)
	if !ok {
		utils.Fatal("IM Hi 解析 FORM 失败")
		return "err"
	}
	echo, ok := form["echostr"]
	// 校验 token
	if ok {
		signature := form["signature"]
		timestamp := form["timestamp"]
		rn := form["rn"]
		if utils.CheckSignature(signature, timestamp, rn) {
			return echo
		}
		utils.Fatal("tocken 校验失败")
		return "err"
	}

	// 响应用户消息
	textMessage, err := utils.ReciveTextMessage(form["message"])
	if err != nil {
		errMsg := form["message"] + "解析失败:" + err.Error()
		utils.Fatal(errMsg)
		return utils.SendTextMessage(errMsg, "cyber_sword")
	}
	user := textMessage.User
	content := textMessage.Content
	utils.Notice(user + ":" + content)
	return utils.SendTextMessage("get it", user)
}

// Book detail of book
func (a *App) Book(params map[string]interface{}) (int, string, interface{}) {
	return utils.OK, "Book ok", params
}

// BookList list of book
func (a *App) BookList(params map[string]interface{}) (int, string, interface{}) {

	return utils.OK, "BookList ok", params
}
