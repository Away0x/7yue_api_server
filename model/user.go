package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"column:name;unique;index:name;not null" json:"name"`
	Key string `gorm:"column:key;unique;index:key;not null" json:"key"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) Create() error {
	return DB.Create(&u).Error
}

// 是否存在该 key
func IsExistKey(key string) error {
	u := &User{Key: key}
	d := DB.Where(u).First(&u)
	if d.Error != nil {
		return d.Error
	}
	return nil
}