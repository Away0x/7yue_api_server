package book

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/handler"
	"github.com/Away0x/7yue_api_server/utils/validate"
	"github.com/Away0x/7yue_api_server/constant/errno"
	"github.com/Away0x/7yue_api_server/model"
	"github.com/Away0x/7yue_api_server/constant"
)

// 获取热门书籍(概要)
func HotList(c *gin.Context) {
	// 1. 响应数据
	handler.SendResponse(c, nil, []interface{}{})
}

// @Summary 获取喜欢书籍数量
// @Description 获取喜欢书籍数量
// @Tags Book
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Success 200 {object} handler.Response
// @Router /v1/book/favor_count [get]
func FavorCount(c *gin.Context) {
	// 1. 获取数据
	key := c.MustGet("appkey").(string)
	_, ids, err := model.GetUserFavorList(key, []int{constant.BOOK_TYPE_CODE})
	if err != nil {
		handler.SendResponse(c, nil, gin.H{
			"count": 0,
		})
		return
	}

	// 2. 响应数据
	handler.SendResponse(c, nil, gin.H{
		"count": len(ids),
	})
}

// @Summary 获取书籍点赞情况
// @Description 获取书籍点赞情况
// @Tags Book
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Param book_id path int true "书籍 id"
// @Success 200 {object} handler.Response
// @Router /v1/book/favor/{book_id} [get]
func FavorStatus(c *gin.Context) {
	// 1. 参数验证
	id, err := validate.NumberParamsValidate(c,"book_id")
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	// 2. 查询数据
	key := c.MustGet("appkey").(string)
	favors, _ := model.GetFavors(id, constant.BOOK_TYPE_CODE)
	is_favor := model.IsFavor(key, uint(id), constant.BOOK_TYPE_CODE, favors)

	// 3. 响应数据
	handler.SendResponse(c, nil, gin.H{
		"fav_nums": len(favors),
		"id": id,
		"like_status": is_favor,
	})
}

// @Summary 获取热搜关键字
// @Description 获取热搜关键字
// @Tags Book
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Success 200 {object} handler.Response
// @Router /v1/book/hot_keyword [get]
func HotKeyword(c *gin.Context) {
	// 1. 获取数据
	words, err := model.GetAllHotKeyWord()
	if err != nil {
		handler.SendResponse(c, nil, gin.H{
			"hot": []interface{}{},
		})
		return
	}

	// 2. 响应数据
	handler.SendResponse(c, nil, gin.H{
		"hot": words,
	})
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
