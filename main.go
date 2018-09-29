package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/router"
	"log"
	"net/http"
)

func main() {
	g := gin.New()

	router.Register(g)

	log.Printf("Start to listening the incoming requests on http address: %s", ":8888")
	log.Fatal(http.ListenAndServe(":8888", g).Error())
}
