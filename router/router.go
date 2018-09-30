package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/router/middleware"
		"github.com/Away0x/7yue_api_server/handler/classic"
	"github.com/Away0x/7yue_api_server/handler"
	"github.com/Away0x/7yue_api_server/constant/errno"
	"github.com/Away0x/7yue_api_server/handler/like"
	"github.com/Away0x/7yue_api_server/handler/book"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/Away0x/7yue_api_server/handler/user"
)

func Register(g *gin.Engine) *gin.Engine {
	// 注册中间件
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(middleware.Options)

	// 注册路由
	// 404
	g.NoRoute(func(c *gin.Context) {
		handler.SendResponse(c, errno.RouterError, nil)
	})
	// swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// index
	g.GET("", user.Index)

	v1 := g.Group("/v1")
	{
		// app key
		userRoute := v1.Group("/user")
		{
			userRoute.GET("", user.List)
			userRoute.POST("/register", user.Register)
		}

		// 期刊
		classicRouter := v1.Group("/classic")
		classicRouter.Use(middleware.KeyAuth)
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
		bookRouter := v1.Group("/book")
		classicRouter.Use(middleware.KeyAuth)
		{
			// 获取热门书籍(概要)
			bookRouter.GET("/hot_list", book.HotList)
			// 获取书籍短评
			bookRouter.GET("/short_comment/:book_id", book.ShortComment)
			// 获取喜欢书籍数量
			bookRouter.GET("/favor_count", book.FavorCount)
			// 获取书籍点赞情况
			bookRouter.GET("/favor/:book_id", book.FavorStatus)
			// 新增短评
			bookRouter.POST("/add/short_comment", book.AddShortComment)
			// 获取热搜关键字
			bookRouter.GET("/hot_keyword", book.HotKeyword)
			// 书籍搜索
			bookRouter.GET("/search", book.Search)
			// 获取书籍详细信息
			bookRouter.GET("/detail/:book_id", book.Detail)
		}

		// 点赞
		likeRouter := v1.Group("/like")
		classicRouter.Use(middleware.KeyAuth)
		{
			// 进行点赞
			likeRouter.POST("", like.Like)
			// 取消点赞
			likeRouter.POST("/cancel", like.Cancel)
		}
	}


	return g
}
