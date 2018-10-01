package mock

import (
	"github.com/Away0x/7yue_api_server/model"
	"fmt"
)

func PushDataIntoUserTable() {
	user := model.User{Name: "wutong", Key: "admin"}

	if err := user.Create(); err != nil {
		fmt.Println("user model 创建失败: " + err.Error())
	}

	user2 := model.User{Name: "wutong2", Key: "admin2"}

	if err := user2.Create(); err != nil {
		fmt.Println("user model 创建失败: " + err.Error())
	}
}