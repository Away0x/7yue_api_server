package mock

import (
	"github.com/Away0x/7yue_api_server/model"
	"fmt"
)

func PushDataIntoBookTable() {
	books := []model.Book{
		{
			Author: "[美]保罗·格雷厄姆",
			BookId: 7,
			Image: "https://img3.doubanio.com/lpic/s4669554.jpg",
			Title: "黑客与画家",
		},
		{
			Author: "MarkPilgrim",
			BookId: 65,
			Image: "https://img3.doubanio.com/lpic/s4059293.jpg",
			Title: "Dive Into Python 3",
		},
		{
			Author: "MagnusLieHetland",
			BookId: 183,
			Image: "https://img3.doubanio.com/lpic/s4387251.jpg",
			Title: "Python基础教程",
		},
		{
			Author: "[哥伦比亚]加西亚·马尔克斯",
			BookId: 1002,
			Image: "https://img3.doubanio.com/lpic/s6384944.jpg",
			Title: "百年孤独",
		},
		{
			Author: "[日]岩井俊二",
			BookId: 1049,
			Image: "https://img1.doubanio.com/view/subject/l/public/s29775868.jpg",
			Title: "情书",
		},
		{
			Author: "[美]乔治·R·R·马丁",
			BookId: 1061,
			Image: "https://img3.doubanio.com/lpic/s1358984.jpg",
			Title: "冰与火之歌（卷一）",
		},
		{
			Author: "[日]东野圭吾",
			BookId: 1120,
			Image: "https://img3.doubanio.com/lpic/s4610502.jpg",
			Title: "白夜行",
		},
		{
			Author: "金庸",
			BookId: 1166,
			Image: "https://img1.doubanio.com/lpic/s23632058.jpg",
			Title: "天龙八部",
		},
		{
			Author: "[日]东野圭吾",
			BookId: 1308,
			Image: "https://img3.doubanio.com/lpic/s3814606.jpg",
			Title: "恶意",
		},
		{
			Author: "[英]J·K·罗琳",
			BookId: 1339,
			Image: "https://img3.doubanio.com/lpic/s1074376.jpg",
			Title: "哈利·波特与阿兹卡班的囚徒",
		},
		{
			Author: "韩寒",
			BookId: 1383,
			Image: "https://img1.doubanio.com/lpic/s3557848.jpg",
			Title: "他的国",
		},
		{
			Author: "[英]J·K·罗琳",
			BookId: 1398,
			Image: "https://img1.doubanio.com/lpic/s2752367.jpg",
			Title: "哈利·波特与死亡圣器",
		},
		{
			Author: "王小波",
			BookId: 1560,
			Image: "https://img1.doubanio.com/lpic/s3463069.jpg",
			Title: "三十而立",
		},
		{
			Author: "[伊朗]玛赞·莎塔碧",
			BookId: 7821,
			Image: "https://img3.doubanio.com/lpic/s6144591.jpg",
			Title: "我在伊朗长大",
		},
		{
			Author: "[日]村上春树",
			BookId: 8854,
			Image: "https://img1.doubanio.com/lpic/s29494718.jpg",
			Title: "远方的鼓声",
		},
		{
			Author: "三毛",
			BookId: 8866,
			Image: "https://img3.doubanio.com/lpic/s2393243.jpg",
			Title: "梦里花落知多少",
		},
		{
			Author: "韩寒",
			BookId: 15198,
			Image: "https://img1.doubanio.com/lpic/s1080179.jpg",
			Title: "像少年啦飞驰",
		},
		{
			Author: "鲁迅",
			BookId: 15984,
			Image: "https://img3.doubanio.com/lpic/s27970504.jpg",
			Title: "朝花夕拾",
		},
		{
			Author: "[日]井上雄彦",
			BookId: 21050,
			Image: "https://img3.doubanio.com/lpic/s2853431.jpg",
			Title: "灌篮高手31",
		},
		{
			Author: "[日]新井一二三",
			BookId: 51664,
			Image: "https://img3.doubanio.com/lpic/s29034294.jpg",
			Title: "东京时味记",
		},
	}

	for _, b := range books {
		if err := b.Create(); err != nil {
			fmt.Println("book model 创建失败: " + err.Error())
		}
	}
}