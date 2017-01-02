package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"strconv"
	"time"
)

// TextMessage recive from baidu hi
type TextMessage struct {
	User       string `xml:"FromUserName"`
	CreateTime int32  `xml:"CreateTime"`
	MsgType    string `xml:"MsgType"`
	Content    string `xml:"Content"`
	// Groups  []string `xml:"Group>Value"`
}

// CheckSignature token
func CheckSignature(signature string, timestamp string, rn string) bool {
	token := "gk2017"
	check := MD5(rn + timestamp + token)
	return check == signature
}

// MD5 get md5 from s
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	c := h.Sum(nil)
	return hex.EncodeToString(c)
}

// SendTextMessage get response string
func SendTextMessage(message string, user string) string {
	ret := "<xml><ToUserName><![CDATA[" + user + "]]></ToUserName>"
	ret += "<CreateTime>" + strconv.FormatInt(time.Now().Unix(), 10) + "</CreateTime>"
	ret += "<MsgType><![CDATA[text]]></MsgType><Content><![CDATA[" + message + "]]></Content></xml>"
	return ret
}

// ReciveTextMessage parse XML
// <xml>
// <FromUserName><![CDATA[fromUser]]></FromUserName>
// <CreateTime></CreateTime>
// <MsgType>text</MsgType>
// <Content><![CDATA[this is a test]]></Content>
// <xml>
func ReciveTextMessage(message string) (TextMessage, error) {
	v := TextMessage{User: "none", MsgType: "text", Content: "none"}
	err := xml.Unmarshal([]byte(message), &v)
	return v, err
}
