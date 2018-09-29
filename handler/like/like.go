package like

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/handler"
	"github.com/Away0x/7yue_api_server/constant/errno"
)

// 进行点赞
func Like(c *gin.Context) {
	handler.SendResponse(c, errno.OK, nil)
}

// 取消点赞
func Cancel(c *gin.Context) {
	handler.SendResponse(c, errno.OK, nil)
}