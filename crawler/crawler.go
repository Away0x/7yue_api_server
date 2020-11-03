package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	BOOK_SEARCH_API = "http://t.talelin.com/v2/book/search?summary=%s&q=%s&start=%s&count=%s"
	BOOK_DETAIL_API = "http://t.talelin.com/v2/book/isbn/%s"
)

func SearchBooks(summary string, q string, start string, count string) interface{} {
	api := fmt.Sprintf(BOOK_SEARCH_API, summary, q, start, count)

	r, err := http.Get(api)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// 解析
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	json_str := strings.TrimSpace(string(b))

	// json
	var res interface{}
	json.Unmarshal([]byte(json_str), &res)

	return res
}

func GetBookDetailByIsbn(isbn string) map[string]interface{} {
	api := fmt.Sprintf(BOOK_DETAIL_API, isbn)

	r, err := http.Get(api)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// 解析
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	json_str := strings.TrimSpace(string(b))

	// json
	var j map[string]interface{}
	json.Unmarshal([]byte(json_str), &j)

	if j["isbn"] == nil || j["isbn"] == "" {
		return nil
	}

	book := map[string]interface{}{
		"author":     j["author"],
		"binding":    j["binding"],
		"category":   j["category"],
		"id":         j["id"],
		"image":      j["image"],
		"images":     j["images"],
		"isbn":       j["isbn"],
		"pages":      j["pages"],
		"price":      j["price"],
		"pubdate":    j["pubdate"],
		"publisher":  j["publisher"],
		"subtitle":   j["subtitle"],
		"summary":    j["summary"],
		"title":      j["title"],
		"translator": j["translator"],
	}

	return book
}
