package model

import "github.com/jinzhu/gorm"

type Favor struct {
	gorm.Model
	UserKey string  `gorm:"column:user_key;not null" json:"user_key"`
	Type int        `gorm:"column:type;not null" json:"type"`
	TargetId int    `gorm:"column:target_id;not null" json:"target_id"`
}

func (Favor) TableName() string {
	return "favor"
}

// 进行点赞
func (f *Favor) Create() error {
	return DB.Create(&f).Error
}

// 取消点赞
func (f *Favor) Delete() error {
	return DB.Delete(&f).Error
}

// 获取点赞
func GetFavor(user_key string, target_id int, _type int) (*Favor, error) {
	f := &Favor{UserKey: user_key, TargetId: target_id, Type: _type}
	d := DB.Where(f).First(&f)
	return f, d.Error
}

// 是否点赞 0 未，1 是
func IsFavor(user_key string, target_id int, _type int) int {
	if _, err := GetFavor(user_key, target_id, _type); err != nil {
		return 0
	}
	return 1
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