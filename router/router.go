package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/router/middleware"
		"github.com/Away0x/7yue_api_server/handler/classic"
	"github.com/Away0x/7yue_api_server/handler"
	"github.com/Away0x/7yue_api_server/constant/errno"
)

func Register(g *gin.Engine) *gin.Engine {
	// 注册中间件
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(middleware.Options)
	g.Use(middleware.KeyAuth)

	// 注册路由
	g.NoRoute(func(c *gin.Context) {
		handler.SendResponse(c, errno.RouterError, nil)
	})
	v1 := g.Group("/v1")
	{
		// 期刊
		classicRouter := v1.Group("/classic")
		{
			// 获取最新一期
			classicRouter.GET("/latest", classic.Latest)
			// 获取当前一期的下一期
			classicRouter.GET("/next/:index", classic.Next)
			// 获取某一期详细信息
			classicRouter.GET("/detail/:type/:id", classic.Detail)
			// 获取当前一期的上一期
			classicRouter.GET("/previous/:index", classic.Previous)
			// 获取点赞信息
			classicRouter.GET("/favor/:type/:id", classic.Like)
			// 获取我喜欢的期刊
			classicRouter.GET("/favor", classic.Favor)
		}

		// 书籍
		//bookRouter := v1.Group("/book")
		//{
		//
		//}

		// 点赞
		//likeRouter := v1.Group("/like")
		//{
		//
		//}
	}


	return g
}
