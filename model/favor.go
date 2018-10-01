package model

import (
	"github.com/jinzhu/gorm"
	"github.com/Away0x/7yue_api_server/constant/errno"
)

type Favor struct {
	gorm.Model
	UserKey string  `gorm:"column:user_key;not null" json:"user_key"`
	Type int        `gorm:"column:type;not null" json:"type"`
	TargetId int    `gorm:"column:target_id;not null" json:"target_id"`
}

func (Favor) TableName() string {
	return "favor"
}

func (f *Favor) Create() error {
	return DB.Create(&f).Error
}

// 点赞
func AddFavor(user_key string, target_id int, _type int) error {
	if _, err := GetFavor(user_key, target_id, _type); err == nil {
		return errno.FavoredError
	}
	f := Favor{UserKey: user_key, TargetId: target_id, Type: _type}
	return f.Create()
}

// 取消点赞
func CancelFavor(user_key string, target_id int, _type int) error {
	if _, err := GetFavor(user_key, target_id, _type); err != nil {
		return errno.NoFavorError
	}
	f := Favor{UserKey: user_key, TargetId: target_id, Type: _type}
	return DB.Where(f).Delete(&Favor{}).Error
}

// 获取点赞
func GetFavor(user_key string, target_id int, _type int) (*Favor, error) {
	f := &Favor{UserKey: user_key, TargetId: target_id, Type: _type}
	d := DB.Where(f).First(&f)
	return f, d.Error
}

// 获取所有用户对该数据的点赞
func GetFavors(target_id int, _type int) ([]*Favor, error) {
	f := make([]*Favor, 0)
	d := DB.Where("target_id = ? AND type = ?", target_id, _type).Find(&f)
	return f, d.Error
}

// 获取所有点赞
func GetAllFavors() ([]*Favor, error) {
	f := make([]*Favor, 0)
	d := DB.Find(&f)
	return f, d.Error
}

// 是否点赞 0 未，1 是
func IsFavor(user_key string, target_id uint, _type int, favors []*Favor) int {
	for _, item := range favors {
		if item.UserKey == user_key && target_id == uint(item.TargetId) && _type == item.Type {
			return 1
		}
	}
	return 0
}

// 获取用户所有的点赞
func GetUserFavorList(user_key string, types []int) ([]*Favor, []int, error) {
	fs := make([]*Favor, 0)
	d := DB.Where("user_key = ? AND type in (?)", user_key, types).Find(&fs)
	if d.Error != nil {
		return nil, nil, d.Error
	}

	ids := make([]int, len(fs))
	for index, f := range fs {
		ids[index] = f.TargetId
	}

	return fs, ids, d.Error
}