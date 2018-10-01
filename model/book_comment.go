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

func (b *BookComment) Create() error {
	return DB.Create(&b).Error
}

func (b *BookComment) AddNums() error {
	num := b.Nums + 1
	return DB.Model(&b).Update("nums", num).Error
}

func GetThisBookAllComments(book_id int) ([]*BookComment, error) {
	comments := make([]*BookComment, 0)
	d := DB.Where("book_id = ?", book_id).Find(&comments)
	return comments, d.Error
}