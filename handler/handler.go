package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Away0x/7yue_api_server/constant/errno"
)

type Response struct {
	Code    int           `json:"error_code"`
	Error   string        `json:"error_message"`
	Data    interface{}  `json:"msg"`
	Request string        `json:"request"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	method := c.Request.Method
	api := c.Request.URL.String()

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Error: message,
		Data:    data,
		Request: method + " " + api,
	})
}
