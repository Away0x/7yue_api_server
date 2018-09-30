package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Away0x/7yue_api_server/handler"
)

// 首页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// 注册
// @Summary 注册
// @Description注册
// @Tags user
// @Accept json
// @Produce json
// @Param user body string true "Create a new user"
// @Success 200 {object} interface{} "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	// 1. 响应数据
	handler.SendResponse(c, nil, "ok")
}

// 列出 key list
func List(c *gin.Context) {
	// 1. 响应数据
	handler.SendResponse(c, nil, "ok")
}