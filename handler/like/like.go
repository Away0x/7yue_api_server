package like

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/handler"
		"github.com/Away0x/7yue_api_server/constant/errno"
	"github.com/Away0x/7yue_api_server/model"
)

// @Summary 进行点赞
// @Description 进行点赞
// @Tags Like
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Param body body like.LikeBody true "点赞的信息"
// @Success 200 {object} handler.Response
// @Router /v1/like [post]
func Like(c *gin.Context) {
	// 1. 参数绑定
	var body LikeBody
	if err := c.ShouldBind(&body); err != nil {
		handler.SendResponse(c, errno.New(errno.JsonError, nil).Add(err.Error()), nil)
		return
	}
	key := c.MustGet("appkey").(string)

	// 2. 数据库操作
	if err := model.AddFavor(key, body.ArtId, body.Type); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	// 2. 响应数据
	handler.SendResponse(c, nil, "ok")
}

// @Summary 取消点赞
// @Description 取消点赞
// @Tags Like
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Param body body like.LikeBody true "取消点赞的信息"
// @Success 200 {object} handler.Response
// @Router /v1/like/cancel [post]
func Cancel(c *gin.Context) {
	// 1. 参数绑定
	var body LikeBody
	if err := c.ShouldBind(&body); err != nil {
		handler.SendResponse(c, errno.New(errno.JsonError, nil).Add(err.Error()), nil)
		return
	}
	key := c.MustGet("appkey").(string)

	// 2. 数据库操作
	if err := model.CancelFavor(key, body.ArtId, body.Type); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	// 2. 响应数据
	handler.SendResponse(c, nil, "ok")
}