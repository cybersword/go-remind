package app

import (
	"database/sql"

	"github.com/cybersword/go-remind/utils"
)

// App entry
type App struct {
	Name string
}

// controllers

// Hi /lab/hi/alpha token is gk2017
func (a *App) Hi(params map[string]interface{}) string {

	form, ok := params["FORM"].(map[string]string)
	if !ok {
		utils.Fatal("IM Hi 解析 FORM 失败", params["FORM"])
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
	// CREATE TABLE `book_list` (
	//   `id`               INT UNSIGNED     AUTO_INCREMENT       COMMENT '自增主键',
	//   `status`           TINYINT UNSIGNED NOT NULL DEFAULT 0   COMMENT '0未购买,1配送中,2已下发',
	//   `book_name`        VARCHAR(512)     NOT NULL DEFAULT ''  COMMENT '书名',
	//   `isbn`             VARCHAR(13)      NOT NULL DEFAULT ''  COMMENT 'ISBN-13',
	//   `url`              VARCHAR(64)      NOT NULL DEFAULT ''  COMMENT 'http://product.dangdang.com/23910258.html',
	//   `user_name`        VARCHAR(32)      NOT NULL DEFAULT ''  COMMENT '姓名',
	//   `price`            FLOAT            NOT NULL DEFAULT 0.0 COMMENT '姓名',
	//   `memo`             VARCHAR(1024)    NOT NULL DEFAULT ''  COMMENT '系统备注',
	//   `create_time`      TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
	//   `update_time`      TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	//   PRIMARY KEY (`id`)
	// )
	//   ENGINE = InnoDB
	//   DEFAULT CHARSET = utf8;
	// INSERT INTO book_list SET status=2, book_name="深入理解Nginx：模块开发与架构解析（第2版）",
	// isbn="9787111526254", url="http://product.dangdang.com/23910258.html", user_name="胡明清", price=78.2;
	// INSERT INTO book_list SET status=2, book_name="运营之光：我的互联网运营方法论与自白",
	// isbn="9787121298097", url="http://product.dangdang.com/24029311.html", user_name="高娟", price=28.9;
	// mysql -hnj02-map-tushang01.nj02 -uroot -proot guoke_lab
	dsnLab := "root:root@tcp(localhost:3308)/guoke_lab?charset=utf8"
	db, err := sql.Open("mysql", dsnLab)
	if err != nil {
		utils.Fatal(err)
		return err.Error()
	}
	defer db.Close()
	if content == "ls" {
		var ls string
		s := "SELECT book_name, user_name FROM book_list"
		rows, err := db.Query(s)
		if err != nil {
			utils.Fatal(err)
			return err.Error()
		}
		var bookName string
		var userName string
		for rows.Next() {
			err = rows.Scan(&bookName, &userName)
			if err != nil {
				utils.Fatal(err)
				return err.Error()
			}
			ls += userName + ": " + bookName + "\n"
		}
		return utils.SendTextMessage(ls, user)
	}

	return utils.SendTextMessage("get it: "+content, user)
}

// Book detail of book
func (a *App) Book(params map[string]interface{}) (int, string, interface{}) {
	return utils.OK, "Book ok", params
}

// BookList list of book
func (a *App) BookList(params map[string]interface{}) (int, string, interface{}) {

	return utils.OK, "BookList ok", params
}
