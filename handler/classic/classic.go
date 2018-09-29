package classic

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/handler"
	"github.com/Away0x/7yue_api_server/constant/errno"
		"github.com/Away0x/7yue_api_server/utils"
)

// 获取最新一期
func Latest(c *gin.Context) {
	handler.SendResponse(c, errno.OK, nil)
}

// 获取当前一期的下一期
func Next(c *gin.Context) {
	// 参数验证
	index, err := validate.NumberParamsValidate(c,"index")
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, errno.OK, index)
}

// 获取某一期详细信息
func Detail(c *gin.Context) {
	// 参数验证
	_type, err := validate.ClassicTypeParamsValidate(c, "type")
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	id, err := validate.NumberParamsValidate(c,"id")
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, gin.H{
		"type": _type,
		"id": id,
	})
}

//获取当前一期的上一期
func Previous(c *gin.Context) {
	// 参数验证
	index, err := validate.NumberParamsValidate(c,"index")
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, errno.OK, index)
}

// 获取点赞信息
func Like(c *gin.Context) {
	// 参数验证
	_type, err := validate.ClassicTypeParamsValidate(c, "type")
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	id, err := validate.NumberParamsValidate(c,"id")
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, gin.H{
		"type": _type,
		"id": id,
	})
}

// 获取我喜欢的期刊
func Favor(c *gin.Context) {
	handler.SendResponse(c, errno.OK, nil)
}
