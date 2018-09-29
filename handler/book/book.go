package book

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/handler"
		"github.com/Away0x/7yue_api_server/utils"
	"github.com/Away0x/7yue_api_server/constant/errno"
)

// 获取热门书籍(概要)
func HotList(c *gin.Context) {
	// 1. 响应数据
	handler.SendResponse(c, nil, "ok")
}

// 获取书籍短评
func ShortComment(c *gin.Context) {
	// 1. 参数验证
	id, err := validate.NumberParamsValidate(c,"book_id")
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	// 2. 响应数据
	handler.SendResponse(c,nil, id)
}

// 获取喜欢书籍数量
func FavorCount(c *gin.Context) {
	// 1. 响应数据
	handler.SendResponse(c, nil, "ok")
}

// 获取书籍点赞情况
func FavorStatus(c *gin.Context) {
	// 1. 参数验证
	id, err := validate.NumberParamsValidate(c,"book_id")
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	// 2. 响应数据
	handler.SendResponse(c, nil, id)
}

// 获取热搜关键字
func HotKeyword(c *gin.Context) {
	// 1. 响应数据
	handler.SendResponse(c, nil, "ok")
}

// 新增短评
func AddShortComment(c *gin.Context) {
	// 1. 参数绑定
	var body AddShortCommentBody
	if err := c.ShouldBind(&body); err != nil {
		handler.SendResponse(c, errno.New(errno.JsonError, nil).Add(err.Error()), nil)
		return
	}

	// 2. 响应数据
	handler.SendResponse(c, nil, body)
}

// 书籍搜索
func Search(c *gin.Context) {
	// 1. 获取参数
	start := c.DefaultQuery("start", "0")     // 开始记录数，默认为 0
	count := c.DefaultQuery("count", "20")    // 记录条数，默认为 20，超过依然按照 20 条计算
	summary := c.DefaultQuery("summary", "0") // 返回完整或简介,默认为 0，0 为完整内容，1 为简介
	q := c.Query("q")                                     // 搜索内容，比如你想搜索 python 相关书籍，则输入 python

	// 2. 响应数据
	handler.SendResponse(c, nil, gin.H{
		"start": start,
		"count": count,
		"summary": summary,
		"q": q,
	})
}

// 获取书籍详细信息
func Detail(c *gin.Context) {
	// 1. 参数验证
	id, err := validate.NumberParamsValidate(c,"book_id")
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	// 2. 响应数据
	handler.SendResponse(c, nil, id)
}
