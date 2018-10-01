package mock

import (
	"github.com/Away0x/7yue_api_server/model"
	"fmt"
)

func PushDataIntoHotKeyWordTable() {
	keywords := []string{
		"Python",
		"哈利·波特",
		"村上春树",
		"东野圭吾",
		"白夜行",
		"韩寒",
		"金庸",
		"王小波",
	}

	for _, word := range keywords {
		h := model.HotKeyword{Text: word}
		if err := h.Create(); err != nil {
			fmt.Println("hot_keyword model 创建失败: " + err.Error())
		}
	}
}