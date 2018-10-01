package like

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/handler"
			"github.com/Away0x/7yue_api_server/constant/errno"
	"github.com/Away0x/7yue_api_server/services"
)

// 进行点赞
func Like(c *gin.Context) {
	// 1. 参数绑定
	var body LikeBody
	if err := c.ShouldBind(&body); err != nil {
		handler.SendResponse(c, errno.New(errno.JsonError, nil).Add(err.Error()), nil)
		return
	}
	key := c.MustGet("appkey").(string)

	// 2. 数据库操作
	if err := services.Like(key, body.ArtId, body.Type); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	// 2. 响应数据
	handler.SendResponse(c, nil, "ok")
}

// 取消点赞
func Cancel(c *gin.Context) {
	// 1. 参数绑定
	var body LikeBody
	if err := c.ShouldBind(&body); err != nil {
		handler.SendResponse(c, errno.New(errno.JsonError, nil).Add(err.Error()), nil)
		return
	}
	key := c.MustGet("appkey").(string)

	// 2. 数据库操作
	if err := services.CancelLike(key, body.ArtId, body.Type); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	// 2. 响应数据
	handler.SendResponse(c, nil, "ok")
}