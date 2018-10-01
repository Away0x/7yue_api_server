package model

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`

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

// 所有用户信息
func GetAllUser() ([]*User, error) {
	us := make([]*User, 0)
	d := DB.Find(&us)
	return us, d.Error
}