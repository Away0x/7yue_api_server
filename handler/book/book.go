package book

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/handler"
	"github.com/Away0x/7yue_api_server/constant/errno"
)

// 获取热门书籍(概要)
func HotList(c *gin.Context) {
	handler.SendResponse(c, errno.OK, nil)
}

// 获取热门书籍(概要)
func ShortComment(c *gin.Context) {
	handler.SendResponse(c, errno.OK, nil)
}

// 获取喜欢书籍数量
func FavorCount(c *gin.Context) {
	handler.SendResponse(c, errno.OK, nil)
}

// 获取书籍点赞情况
func FavorStatus(c *gin.Context) {
	handler.SendResponse(c, errno.OK, nil)
}

// 新增短评
func AddShortComment(c *gin.Context) {
	handler.SendResponse(c, errno.OK, nil)
}

// 获取热搜关键字
func HotKeyword(c *gin.Context) {
	handler.SendResponse(c, errno.OK, nil)
}

// 书籍搜索
func Search(c *gin.Context) {
	handler.SendResponse(c, errno.OK, nil)
}

// 获取书籍详细信息
func Detail(c *gin.Context) {
	handler.SendResponse(c, errno.OK, nil)
}
