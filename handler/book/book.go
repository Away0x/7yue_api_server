package book

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/handler"
	"github.com/Away0x/7yue_api_server/utils/validate"
	"github.com/Away0x/7yue_api_server/constant/errno"
	"github.com/Away0x/7yue_api_server/model"
	"github.com/Away0x/7yue_api_server/constant"
	"github.com/Away0x/7yue_api_server/crawler"
)

// @Summary 获取热门书籍(概要)
// @Description获取热门书籍(概要)
// @Tags Book
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Success 200 {array} model.BookSerializer
// @Router /v1/book/hot_list [get]
func HotList(c *gin.Context) {
	// 1. 获取数据
	books, err := model.GetAllBook()
	if err != nil {
		handler.SendResponse(c, nil, []interface{}{})
		return
	}
	key := c.MustGet("appkey").(string)
	favors, _ := model.GetAllFavors()

	// 2. 数据整理
	result := make([]model.BookSerializer, len(books))         // 最终响应的数据
	cur_user_favor_status_map := make(map[int]int, len(books)) // 当前用户对于各本书的点赞状态
	all_user_favor_nums_map := make(map[int]int, len(books))   // 所有用户对于各本书的点赞数

	for _, favor := range favors {
		// 当前用户是否点过赞
		if favor.UserKey == key && favor.Type == constant.BOOK_TYPE_CODE {
			cur_user_favor_status_map[favor.TargetId] = 1
		}
		// 用户点赞数统计
		if favor.Type == constant.BOOK_TYPE_CODE {
			all_user_favor_nums_map[favor.TargetId] = all_user_favor_nums_map[favor.TargetId] + 1
		}
	}

	for index, book := range books {
		like_status := 0
		if cur_user_favor_status_map[book.BookId] == 1 {
			like_status = 1
		}
		result[index] = book.Serializer(like_status, all_user_favor_nums_map[book.BookId])
	}


	// 3. 响应数据
	handler.SendResponse(c, nil, result)
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

// @Summary 新增短评
// @Description 新增短评
// @Tags Book
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Param body body book.AddShortCommentBody true "短评信息"
// @Success 200 {object} handler.Response
// @Router /v1/book/add/short_comment [post]
func AddShortComment(c *gin.Context) {
	// 1. 参数绑定
	var body AddShortCommentBody
	if err := c.ShouldBind(&body); err != nil {
		handler.SendResponse(c, errno.New(errno.JsonError, nil).Add(err.Error()), nil)
		return
	}

	// 2. 操作数据库 (短评已存在则 nums +1，否则创建)
	if comment, err := model.IsExistThisBookComments(body.BookId, body.Content); err == nil {
		if err := comment.AddNums(); err != nil {
			handler.SendResponse(c, err, nil)
			return
		}
	} else {
		new_comment := model.BookComment{Content: body.Content, BookId: body.BookId, Nums: 1}
		if err := new_comment.Create(); err != nil {
			handler.SendResponse(c, err, nil)
			return
		}
	}


	// 2. 响应数据
	handler.SendResponse(c, nil, "ok")
}

// @Summary 获取书籍短评
// @Description 获取书籍短评
// @Tags Book
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Param book_id path string true "书籍 id"
// @Success 200 {object} handler.Response
// @Router /v1/book/short_comment/{book_id} [get]
func ShortComment(c *gin.Context) {
	// 1. 参数验证
	id, err := validate.NumberParamsValidate(c,"book_id")
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	// 2. 获取数据
	comments, _ := model.GetThisBookAllComments(id)

	// 2. 响应数据
	handler.SendResponse(c,nil, gin.H{
		"book_id": id,
		"comments": comments,
	})
}

// @Summary 书籍搜索
// @Description 书籍搜索
// @Tags Book
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Param q query string true "搜索内容,比如你想搜索python相关书籍,则输入python"
// @Param summary query string false "返回完整或简介,默认为0,0为完整内容,1为简介"
// @Param start query string false "开始记录数，默认为0"
// @Param count query string false "记录条数，默认为20,超过依然按照20条计算"
// @Success 200 {object} handler.Response
// @Router /v1/book/search [get]
func Search(c *gin.Context) {
	// 1. 获取参数
	start := c.DefaultQuery("start", "0")     // 开始记录数，默认为 0
	count := c.DefaultQuery("count", "20")    // 记录条数，默认为 20，超过依然按照 20 条计算
	summary := c.DefaultQuery("summary", "0") // 返回完整或简介,默认为 0，0 为完整内容，1 为简介
	q := c.Query("q")                                     // 搜索内容，比如你想搜索 python 相关书籍，则输入 python

	// 2. 调用 api
	res := crawler.SearchBooks(summary, q, start, count)

	// 2. 响应数据
	handler.SendResponse(c, nil, res)
}

// @Summary 通过 isbn 获取书籍详细信息
// @Description 通过 isbn 获取书籍详细信息
// @Tags Book
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Param isbn path string true "书籍的 isbn 号"
// @Success 200 {object} handler.Response
// @Router /v1/book/detail/{isbn} [get]
func Detail(c *gin.Context) {
	// 1. 参数验证
	isbn := c.Param("isbn")
	if isbn == "" {
		handler.SendResponse(c, errno.ParamsError, nil)
		return
	}

	// 2. 调用 api
	res := crawler.GetBookDetailByIsbn(isbn)

	// 2. 响应数据
	handler.SendResponse(c, nil, res)
}
