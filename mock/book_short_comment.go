package mock

import (
	"fmt"
	"github.com/Away0x/7yue_api_server/model"
)

func PushDataIntoBookCommentTable() {
	comments := []model.BookComment{
		// book id 7
		{
			BookId: 7,
			Content: "程序员也有艺术的人生",
			Nums: 1001,
		},
		{
			BookId: 7,
			Content: "黑客and",
			Nums: 475,
		},
		{
			BookId: 7,
			Content: "七月老师好",
			Nums: 107,
		},
		{
			BookId: 7,
			Content: "a",
			Nums: 28,
		},
		{
			BookId: 7,
			Content: "一二三四五六七八九十十一",
			Nums: 26,
		},
		{
			BookId: 7,
			Content: "666",
			Nums: 10,
		},
		{
			BookId: 7,
			Content: "好好的",
			Nums: 8,
		},
		{
			BookId: 7,
			Content: "艺术",
			Nums: 7,
		},
		{
			BookId: 7,
			Content: "555",
			Nums: 6,
		},
		{
			BookId: 7,
			Content: "ABC",
			Nums: 5,
		},
		// book id 65
		{
			BookId: 65,
			Content: "来个沙发！",
			Nums: 478,
		},
		{
			BookId: 65,
			Content: "888",
			Nums: 167,
		},
		{
			BookId: 65,
			Content: "测试一下",
			Nums: 102,
		},
		{
			BookId: 65,
			Content: "444",
			Nums: 2,
		},
		{
			BookId: 65,
			Content: "我喜欢Python很久了",
			Nums: 2,
		},
		{
			BookId: 65,
			Content: "a",
			Nums: 2,
		},
		{
			BookId: 65,
			Content: "测试",
			Nums: 2,
		},
		{
			BookId: 65,
			Content: "刑警本色",
			Nums: 1,
		},
		{
			BookId: 65,
			Content: "qwewqe",
			Nums: 1,
		},
		{
			BookId: 65,
			Content: "你的",
			Nums: 1,
		},
		// book id 183
		{
			BookId: 183,
			Content: "人生苦短，我用Py",
			Nums: 292,
		},
		{
			BookId: 183,
			Content: "I like Py",
			Nums: 113,
		},
		{
			BookId: 183,
			Content: "life",
			Nums: 53,
		},
		{
			BookId: 183,
			Content: "lifeisshort",
			Nums: 2,
		},
		{
			BookId: 183,
			Content: "你好，七月",
			Nums: 2,
		},
		{
			BookId: 183,
			Content: "七月，好python ",
			Nums: 1,
		},
		{
			BookId: 183,
			Content: "经典致敬",
			Nums: 1,
		},
		{
			BookId: 183,
			Content: "还不错",
			Nums: 1,
		},
		{
			BookId: 183,
			Content: "cool",
			Nums: 1,
		},
		{
			BookId: 183,
			Content: "公司",
			Nums: 1,
		},
		// book id 1002
		{
			BookId: 1002,
			Content: "魔幻第一书",
			Nums: 173,
		},
		{
			BookId: 1002,
			Content: "永恒的经典",
			Nums: 171,
		},
		{
			BookId: 1002,
			Content: "一百年，一世纪",
			Nums: 39,
		},
		{
			BookId: 1002,
			Content: "Nice",
			Nums: 1,
		},
		{
			BookId: 1002,
			Content: "我很喜欢",
			Nums: 1,
		},
		{
			BookId: 1002,
			Content: "推图",
			Nums: 1,
		},
		{
			BookId: 1002,
			Content: "啦咯啦咯啦咯啦咯啦咯啦",
			Nums: 1,
		},
		{
			BookId: 1002,
			Content: "凝聚态",
			Nums: 1,
		},
		{
			BookId: 1002,
			Content: "怎么了？",
			Nums: 1,
		},
		{
			BookId: 1002,
			Content: "我写作业了",
			Nums: 1,
		},
		// book id 1049
		{
			BookId: 1049,
			Content: "致敬经典",
			Nums: 109,
		},
		{
			BookId: 1049,
			Content: "天亦老",
			Nums: 60,
		},
		{
			BookId: 1049,
			Content: "人生亦相逢",
			Nums: 26,
		},
		{
			BookId: 1049,
			Content: "爱而不得",
			Nums: 3,
		},
		{
			BookId: 1049,
			Content: "眼泪成诗",
			Nums: 2,
		},
		{
			BookId: 1049,
			Content: "程序员",
			Nums: 2,
		},
		{
			BookId: 1049,
			Content: "大师经典之作",
			Nums: 1,
		},
		{
			BookId: 1049,
			Content: "123",
			Nums: 1,
		},
		{
			BookId: 1049,
			Content: "岩井俊二",
			Nums: 1,
		},
		{
			BookId: 1049,
			Content: "大师",
			Nums: 1,
		},
	}

	for _, c := range comments {
		if err := c.Create(); err != nil {
			fmt.Println("book short comment model 创建失败: " + err.Error())
		}
	}
}
