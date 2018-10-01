package services

import (
	"github.com/Away0x/7yue_api_server/constant"
	"github.com/Away0x/7yue_api_server/model"
	"github.com/Away0x/7yue_api_server/constant/errno"
)

// 点赞
func Like(key string, id int, _type int) error {
	// 要点赞的数据不存在
	var err error
	var target interface{}
	if _type == constant.BOOK_TYPE_CODE {

	} else {
		if target, err = model.GetClassicDetail(id, _type); err != nil {
			return errno.NoClassicError
		}
	}
	// 已点过赞
	if _, err = model.GetFavor(key, id, _type); err == nil {
		return errno.FavoredError
	}

	// 点赞操作
	favor := model.Favor{UserKey: key, TargetId: id, Type: _type}
	if err = favor.Create(); err != nil {
		return err
	}
	// 点赞的是 book 还是 classic，更新点赞数
	switch t := target.(type) {
	case *model.Book:
	case *model.Classic:
		if err = t.FavorClassic(); err != nil {
			return err
		}
	}

	return nil
}

// 取消点赞
func CancelLike(key string, id int, _type int) error {
	// 要取消点赞的数据不存在
	var err error
	var target interface{}
	if _type == constant.BOOK_TYPE_CODE {

	} else {
		if target, err = model.GetClassicDetail(id, _type); err != nil {
			return errno.NoClassicError
		}
	}
	// 没点过赞
	if _, err = model.GetFavor(key, id, _type); err != nil {
		return errno.NoFavorError
	}

	// 取消点赞操作
	favor := model.Favor{UserKey: key, TargetId: id, Type: _type}
	if err = favor.Delete(); err != nil {
		return err
	}
	// 取消点赞的是 book 还是 classic，更新点赞数
	switch t := target.(type) {
	case *model.Book:
	case *model.Classic:
		if err = t.CancelFavorClassic(); err != nil {
			return err
		}
	}

	return nil
}
