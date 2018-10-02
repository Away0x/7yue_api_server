package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/router"
	"log"
	"net/http"
	"github.com/Away0x/7yue_api_server/model"
	"github.com/spf13/viper"
	"github.com/Away0x/7yue_api_server/mock"
	"github.com/Away0x/7yue_api_server/config"
)

func main() {
	config.InitConfig()

	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()
	g.Static("/static", "static")
	g.LoadHTMLGlob("templates/*")
	g.StaticFile("/favicon.ico", "static/favicon.ico")

	db := model.InitDB()
	db.AutoMigrate(
		&model.Classic{},
		&model.Book{},
		&model.BookComment{},
		&model.Favor{},
		&model.User{},
		&model.HotKeyword{})
	defer db.Close()

	// 首次运行时可通过该方法添加一些 mock 数据
	// mockInit()

	router.Register(g)

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("port"))
	log.Fatal(http.ListenAndServe(viper.GetString("port"), g).Error())
}

// mock 数据
func mockInit() {
	// mock user
	mock.PushDataIntoUserTable()

	// mock classic
	mock.PushDataIntoClassicTable()

	// mock book
	mock.PushDataIntoBookTable()

	// mock book short comment
	mock.PushDataIntoBookCommentTable()

	// mock hot keyword
	mock.PushDataIntoHotKeyWordTable()
}