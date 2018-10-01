// 书籍短评
package model

import "github.com/jinzhu/gorm"

type BookComment struct {
	gorm.Model
	Content string `gorm:"varchar(255);column:content;not null"`
	BookId int `gorm:"integer;column:book_id;unique;index:book_id;not null"`
	Nums int `gorm:"column:nums;not null;default:0"`
}

func (BookComment) TableName() string {
	return "book_comment"
}