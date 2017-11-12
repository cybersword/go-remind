package app

import (
	"database/sql"

	"strings"

	"strconv"

	"github.com/cybersword/go-remind/utils"
	"github.com/monnand/goredis"
)

// import "github.com/monnand/goredis"
// import "github.com/garyburd/redigo/redis"

// controllers

// Lab ...
type Lab struct {
	AppName string
}

// Name ...
func (app *Lab) Name() string {
	return "Hi"
}

// Hi /lab/hi/alpha token is gk2017
func (a *Lab) Hi(params map[string]interface{}) string {

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
	// dsnLab := "root:root@tcp(localhost:3308)/test_lab?charset=utf8"
	dsnLab := "dev@tcp(127.0.0.1:3306)/test?charset=utf8"
	db, err := sql.Open("mysql", dsnLab)
	if err != nil {
		utils.Fatal(err)
		return err.Error()
	}
	defer db.Close()

	if content == "ls" {
		// 列清单
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
			ls += "【" + userName + "】—— " + bookName + "\n"
		}
		return utils.SendTextMessage(ls, user)
	} else if strings.Contains(content, "http://product.dangdang.com/") {
		// 买书
		arr := strings.Split(content, " ")
		userName := arr[0]
		book, err := utils.DangDang(arr[1])
		if err != nil {
			utils.Fatal(err)
			return err.Error()
		}
		sql := "INSERT INTO book_list SET book_name='" + book.Name + "', price=" + strconv.FormatFloat(book.Price, 'f', -1, 64)
		sql += ", user_name='" + userName + "', isbn='" + book.ISBN + "', url='" + book.URL + "', memo='" + userName + "'"
		result, err := db.Exec(sql)
		lastID, err := result.LastInsertId()
		utils.Notice(book, lastID)
		return utils.SendTextMessage("ok", user)

	}

	return utils.SendTextMessage("get it: "+content, user)
}

// Book detail of book
func (a *Lab) Book(params map[string]interface{}) (int, string, interface{}) {
	form, ok := params["FORM"].(map[string]string)
	if !ok {
		utils.Fatal("解析FORM失败", params["FORM"])
		return utils.ERROR, "解析FORM失败", params["FORM"]
	}

	bookName := form["book_name"]
	bookURL := form["book_url"]
	bookISBN := form["book_isbn"]
	bookPrice := form["book_price"]
	book := make(map[string]string)
	book["name"] = bookName
	book["url"] = bookURL
	book["isbn"] = bookISBN
	book["price"] = bookPrice

	var client goredis.Client
	// 设置端口为redis默认端口
	client.Addr = "localhost:8379"

	err := client.Hmset("book:"+bookISBN, book)
	if err != nil {
		return utils.ERROR, err.Error(), book
	}

	bName, err := client.Hget("book:"+bookISBN, "name")
	if err != nil {
		return utils.ERROR, err.Error(), book
	}

	// //字符串操作
	// client.Set("a", []byte("hello"))
	// val, _ := client.Get("a")
	// fmt.Println(string(val))
	// client.Del("a")

	// //list操作
	// vals := []string{"a", "b", "c", "d", "e"}
	// for _, v := range vals {
	// 	client.Rpush("l", []byte(v))
	// }
	// dbvals, _ := client.Lrange("l", 0, 4)
	// for i, v := range dbvals {
	// 	println(i, ":", string(v))
	// }
	// client.Del("l")
	return utils.OK, "Book ok", string(bName)
}

// BookList list of book
func (a *Lab) BookList(params map[string]interface{}) (int, string, interface{}) {

	return utils.OK, "BookList ok", params
}
