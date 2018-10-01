package model

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Author  string `gorm:"type:varchar(128);not null;column:author"`
	FavNums int    `gorm:"type:integer;not null;column:fav_nums;default:0"`
	Image   string `gorm:"type:varchar(128);not null;column:image"`
	Title   string `gorm:"type:varchar(128);not null;column:title"`
}

func (Book) TableName() string {
	return "book"
}
