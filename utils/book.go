package utils

// book.go - functions to parse book shop pages
// TODO: 之后整理成 book 类

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
)

// Book Website is "dangdang.com"
type Book struct {
	ISBN    string
	Name    string
	Price   float64
	URL     string
	Website string
}

// DangDang parse dangdang.com
func DangDang(bookURL string) (book Book, err error) {
	book = Book{"", "", 0.0, bookURL, "dangdang.com"}
	doc, err := goquery.NewDocument(bookURL)
	if err != nil {
		return
	}
	// 图书展示区 div.product_main.clearfix div.show_info
	main := doc.Find("div.product_main.clearfix div.show_info")
	name := main.Find("div.name_info h1").Text()
	name = mahonia.NewDecoder("gbk").ConvertString(name)
	book.Name = strings.TrimSpace(name)
	// <p id="dd-price"><span class="yen">¥</span>78.20</p>
	price := main.Find("div.price_info p#dd-price").Text()
	price = strings.TrimSpace(price)
	price = strings.TrimLeft(price, "¥")
	book.Price, err = strconv.ParseFloat(price, 64)

	doc.Find("#product_tab #detail_all #detail_describe ul li").Each(func(i int, s *goquery.Selection) {
		t := mahonia.NewDecoder("gbk").ConvertString(s.Text())
		if strings.Contains(t, "ISBN") {
			book.ISBN = strings.TrimSpace(t)
			book.ISBN = strings.Split(t, "：")[1]
		}
	})

	return
}
