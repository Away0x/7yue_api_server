package model

import "time"

type User struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`

	Name string `gorm:"column:name;unique;index:name;not null" json:"name"`
	Key string `gorm:"column:key;unique;index:key;not null" json:"key"`
}

func (User) TableName() string {
	return "user"
}