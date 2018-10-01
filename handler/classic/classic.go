package classic

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/handler"
	"github.com/Away0x/7yue_api_server/constant/errno"
	"github.com/Away0x/7yue_api_server/utils"
	"github.com/Away0x/7yue_api_server/model"
	"github.com/Away0x/7yue_api_server/utils/validate"
	"github.com/Away0x/7yue_api_server/constant"
)

// @Summary 获取最新一期期刊
// @Description 获取最新一期期刊
// @Tags Classic
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Success 200 {object} model.ClassicSerializer
// @Router /v1/classic/latest [get]
func Latest(c *gin.Context) {
	// 1. 查询数据
	classic, err := model.LatestClassic()
	if err != nil {
		handler.SendResponse(c, errno.NoClassicError, nil)
		return
	}
	key := c.MustGet("appkey").(string)
	favors, _ := model.GetFavors(int(classic.ID), classic.Type)
	is_favor := model.IsFavor(key, classic.ID, classic.Type, favors)

	// 3. 响应数据
	handler.SendResponse(c, errno.OK, classic.Serializer(utils.GetPath(c), is_favor, len(favors)))
}

// @Summary 获取当前一期的下一期期刊
// @Description 获取当前一期的下一期期刊
// @Tags Classic
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Param index path int true "当前期刊的 index"
// @Success 200 {object} model.ClassicSerializer
// @Router /v1/classic/next/{index} [get]
func Next(c *gin.Context) {
	// 1. 参数验证
	index, err := validate.NumberParamsValidate(c,"index")
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	// 2. 查询数据
	classic, err := model.FindNextClassic(index)
	if err != nil {
		handler.SendResponse(c, errno.NoClassicError, nil)
		return
	}
	key := c.MustGet("appkey").(string)
	favors, _ := model.GetFavors(int(classic.ID), classic.Type)
	is_favor := model.IsFavor(key, classic.ID, classic.Type, favors)

	// 3. 响应数据
	handler.SendResponse(c, errno.OK, classic.Serializer(utils.GetPath(c), is_favor, len(favors)))
}

// @Summary 获取当前一期的上一期期刊
// @Description 获取当前一期的上一期期刊
// @Tags Classic
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Param index path int true "当前期刊的 index"
// @Success 200 {object} model.ClassicSerializer
// @Router /v1/classic/previous/{index} [get]
func Previous(c *gin.Context) {
	// 1. 参数验证
	index, err := validate.NumberParamsValidate(c,"index")
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	// 2. 查询数据
	classic, err := model.FindPreviousClassic(index)
	if err != nil {
		handler.SendResponse(c, errno.NoClassicError, nil)
		return
	}
	key := c.MustGet("appkey").(string)
	favors, _ := model.GetFavors(int(classic.ID), classic.Type)
	is_favor := model.IsFavor(key, classic.ID, classic.Type, favors)

	// 3. 响应数据
	handler.SendResponse(c, errno.OK, classic.Serializer(utils.GetPath(c), is_favor, len(favors)))
}

// @Summary 获取某一期详细信息
// @Description 获取某一期详细信息
// @Tags Classic
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Param type path int true "期刊的类型 100 200 300"
// @Param id path int true "当前期刊的 id"
// @Success 200 {object} model.ClassicSerializer
// @Router /v1/classic/detail/{type}/{id} [get]
func Detail(c *gin.Context) {
	// 1. 参数验证
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

	// 2. 查询数据
	classic, err := model.GetClassicDetail(id, _type)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	key := c.MustGet("appkey").(string)
	favors, _ := model.GetFavors(int(classic.ID), classic.Type)
	is_favor := model.IsFavor(key, classic.ID, classic.Type, favors)

	// 3. 响应数据
	handler.SendResponse(c, errno.OK, classic.Serializer(utils.GetPath(c), is_favor, len(favors)))
}

// @Summary 获取点赞信息
// @Description 获取点赞信息
// @Tags Classic
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Param type path int true "期刊的类型 100 200 300"
// @Param id path int true "当前期刊的 id"
// @Success 200 {object} handler.Response
// @Router /v1/classic/favor/{type}/{id} [get]
func Like(c *gin.Context) {
	// 1. 参数验证
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

	// 2. 查询数据
	classic, err := model.GetClassicDetail(id, _type)
	if err != nil {
		handler.SendResponse(c, errno.NoClassicError, nil)
		return
	}
	key := c.MustGet("appkey").(string)
	favors, _ := model.GetFavors(int(classic.ID), classic.Type)
	is_favor := model.IsFavor(key, classic.ID, classic.Type, favors)

	// 3. 响应数据
	handler.SendResponse(c, nil, gin.H{
		"fav_nums": len(favors),
		"id": classic.ID,
		"like_status": is_favor,
	})
}

// @Summary 获取我喜欢的期刊
// @Description 获取我喜欢的期刊
// @Tags Classic
// @Accept  json
// @Produce  json
// @Param appkey query string true "令牌: 测试可用 admin 或你自己的 appkey"
// @Success 200 {array} model.ClassicSerializer
// @Router /v1/classic/favor [get]
func Favor(c *gin.Context) {
	// 1. 获取参数
	start, err := validate.NumberQueryValidate(c,"start", "1")
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	count, err := validate.NumberQueryValidate(c,"count", "20")
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	if count > 20 {
		count = 20
	}
	key := c.MustGet("appkey").(string)

	// 2. 查询数据
	classices, err := model.GetFavorClassices(key, start, count)
	if err != nil {
		handler.SendResponse(c, errno.NoClassicError, [...]interface{}{})
		return
	}
	_, ids, _ := model.GetUserFavorList(key, []int{constant.ClASSIC_TYPE_MOVIE, constant.ClASSIC_TYPE_MUSIC, constant.ClASSIC_TYPE_SENTENCE})
	fav_nums_map := make(map[int]int, len(classices))
	for _, id := range ids {
		fav_nums_map[id] = fav_nums_map[id] + 1
	}

	// 3. 响应数据
	result := make([]model.ClassicSerializer, len(classices))
	for index, classic := range classices {
		result[index] = classic.Serializer(utils.GetPath(c), 1, fav_nums_map[int(classic.ID)])
	}
	handler.SendResponse(c, nil, result)
}
