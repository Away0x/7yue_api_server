package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Away0x/7yue_api_server/handler"
	"golang.org/x/crypto/bcrypt"
	"github.com/Away0x/7yue_api_server/model"
	"github.com/Away0x/7yue_api_server/constant/errno"
)

// 首页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// @Summary 注册
// @Description 注册
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body user.UserBody true "用户名"
// @Success 200 {object} handler.Response
// @Router /v1/user [post]
func Register(c *gin.Context) {
	// 1. 获取数据
	var form UserBody
	if err := c.ShouldBind(&form); err != nil {
		handler.SendResponse(c, errno.New(errno.JsonError, nil).Add(err.Error()), nil)
		return
	}
	if form.UserName == "" {
		handler.SendResponse(c, errno.UserCreateError, nil)
		return
	}
	key, err := bcrypt.GenerateFromPassword([]byte(form.UserName), bcrypt.DefaultCost)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	// 2. 操作数据库
	user := model.User{Name: form.UserName, Key: string(key)}
	if err := user.Create(); err != nil {
		handler.SendResponse(c, errno.UserCreateError, nil)
		return
	}

	// 3. 响应数据
	handler.SendResponse(c, nil, gin.H{
		"name": user.Name,
		"id": user.ID,
		"key": user.Key,
	})
}

// @Summary 获取所有用户
// @Description 获取所有用户
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {array} model.User
// @Router /v1/user [get]
func List(c *gin.Context) {
	// 1. 获取数据
	users, _ := model.GetAllUser()

	// 2. 响应数据
	handler.SendResponse(c, nil, users)
}