package model

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	BookId  int    `gorm:"type:integer;not null;column:book_id"`
	Author  string `gorm:"type:varchar(128);not null;column:author"`
	Image   string `gorm:"type:varchar(128);not null;column:image"`
	Title   string `gorm:"type:varchar(128);not null;column:title"`
}

func (Book) TableName() string {
	return "book"
}

func (b *Book) Create() error {
	return DB.Create(&b).Error
}

func GetBookDetail(book_id int) (*Book, error) {
	b := &Book{}
	d := DB.First(&b, book_id)
	return b, d.Error
}
