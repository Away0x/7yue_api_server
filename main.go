package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/router"
	"log"
	"net/http"
	"github.com/Away0x/7yue_api_server/model"
	"github.com/Away0x/7yue_api_server/config"
)

func main() {
	g := gin.New()
	g.LoadHTMLGlob("templates/*")

	config.InitConfig()

	db := model.InitDB()
	db.AutoMigrate(
		&model.Classic{},
		&model.Book{},
		&model.BookComment{},
		&model.Favor{},
		&model.User{},
		&model.HotKeyword{})
	defer db.Close()

	router.Register(g)

	log.Printf("Start to listening the incoming requests on http address: %s", ":8888")
	log.Fatal(http.ListenAndServe(":8888", g).Error())
}
