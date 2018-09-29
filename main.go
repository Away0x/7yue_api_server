package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/router"
	"log"
	"net/http"
	"github.com/Away0x/7yue_api_server/model"
)

func main() {
	g := gin.New()

	db := model.InitDB()
	defer db.Close()

	router.Register(g)

	log.Printf("Start to listening the incoming requests on http address: %s", ":8888")
	log.Fatal(http.ListenAndServe(":8888", g).Error())
}
