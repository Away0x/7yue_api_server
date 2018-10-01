package mock

import (
	"github.com/Away0x/7yue_api_server/model"
	"fmt"
)

func PushDataIntoClassicTable() {
	classics := []model.Classic{
		{
			Content: "谁念过 千字文章 秋收冬已藏",
			Image: "music.7.png",
			Index: 7,
			Title: "不才 《参商》",
			Type: 200,
			Url: "http://music.163.com/song/media/outer/url?id=29719651.mp3",
		},
		{
			Content: "心上无垢，林间有风",
			Image: "sentence.6.png",
			Index: 6,
			Title: "未名",
			Type: 300,
		},
		{
			Content: "许多人来来去去，相聚又别离",
			Image: "music.5.png",
			Index: 5,
			Title: "好妹妹 《一个人的北京》",
			Type: 200,
			Url: "http://music.163.com/song/media/outer/url?id=26427662.mp3",
		},
		{
			Content: "在变换的生命里，岁月原来是最大的小偷",
			Image: "movie.4.png",
			Index: 4,
			Title: "罗启锐《岁月神偷》",
			Type: 100,
		},
		{
			Content: "你陪我步入蝉夏 越过城市喧嚣",
			Image: "music.1.png",
			Index: 3,
			Title: "花粥 《纸短情长》",
			Type: 200,
			Url: "http://music.163.com/song/media/outer/url?id=557581284.mp3",
		},
		{
			Content: "这个夏天又是一个毕业季",
			Image: "sentence.2.png",
			Index: 2,
			Title: "未名",
			Type: 300,
		},
		{
			Content: "岁月长，衣裳薄",
			Image: "music.3.png",
			Index: 1,
			Title: "杨千嬅《再见二丁目》",
			Type: 200,
			Url: "http://music.163.com/song/media/outer/url?id=317245.mp3",
		},
		{
			Content: "人生不能像做菜，把所有的料准备好才下锅",
			Image: "movie.8.png",
			Index: 8,
			Title: "李安《饮食男女 》",
			Type: 100,
		},
	}

	for _, c := range classics {
		if err := c.Create(); err != nil {
			fmt.Println("classic model 创建失败: " + err.Error())
		}
	}
}