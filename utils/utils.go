package utils

import (
	"github.com/gin-gonic/gin"
)

func GetPath(c *gin.Context) string {
	return "http://" + c.Request.Host
}

