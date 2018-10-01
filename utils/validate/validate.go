package validate

import (
	"strconv"
	"github.com/Away0x/7yue_api_server/constant"
	"github.com/Away0x/7yue_api_server/constant/errno"
	"github.com/gin-gonic/gin"
)

// 参数必须为数字
func NumberParamsValidate(c *gin.Context, key string) (int, error) {
	num, err := strconv.Atoi(c.Param(key))
	if err != nil {
		return -1, errno.New(errno.ParamsError, nil).Addf( "参数 %s 必须为数字类型", key)
	}

	return num, nil
}
func NumberQueryValidate(c *gin.Context, key string, default_value string) (int, error) {
	num, err := strconv.Atoi(c.DefaultQuery(key, default_value))
	if err != nil {
		return -1, errno.New(errno.ParamsError, nil).Addf( "参数 %s 必须为数字类型", key)
	}

	return num, nil
}

// 期刊类型必须为数字，且在取值范围内
func ClassicTypeParamsValidate(c *gin.Context, key string) (int, error) {
	_type, err := strconv.Atoi(c.Param(key))
	if err != nil || (_type != constant.ClASSIC_TYPE_MOVIE && _type != constant.ClASSIC_TYPE_MUSIC && _type != constant.ClASSIC_TYPE_SENTENCE) {
		return -1, errno.New(errno.ParamsError, nil).Addf("type 必须为数字类型，且取值范围为 %d %d %d", constant.ClASSIC_TYPE_MOVIE, constant.ClASSIC_TYPE_MUSIC, constant.ClASSIC_TYPE_SENTENCE)
	}

	return _type, nil
}