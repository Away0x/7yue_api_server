package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/model"
	"github.com/Away0x/7yue_api_server/constant/errno"
	"github.com/Away0x/7yue_api_server/handler"
)

// 跨域
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

// appkey
// header 携带的优先级更高
func KeyAuth(c *gin.Context) {
	key := c.Request.Header.Get("appkey")
	if key == "" {
		key = c.Query("appkey")
	}

	// 没有 key
	if err := model.IsExistKey(key); err != nil || key == "" {
		handler.SendResponse(c, errno.New(errno.AuthError, nil).Addf("appkey 错误 - %s 该用户还未注册", key), nil)
		c.Abort()
		return
	}

	c.Set("appkey", key)
	c.Next()
}