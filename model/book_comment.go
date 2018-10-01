// 书籍短评
package model

import (
	"time"
)

type BookComment struct {
	ID        uint `gorm:"primary_key" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`

	Content string `gorm:"varchar(255);column:content;not null" json:"content"`
	BookId int `gorm:"integer;column:book_id;index:book_id;not null" json:"-"`
	Nums int `gorm:"column:nums;not null;default:1" json:"nums"`
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

func IsExistThisBookComments(book_id int, comment string) (*BookComment, error) {
	c := &BookComment{BookId: book_id, Content: comment}
	d := DB.Where(c).First(&c)
	return c, d.Error
}

func GetThisBookAllComments(book_id int) ([]*BookComment, error) {
	comments := make([]*BookComment, 0)
	d := DB.Where("book_id = ?", book_id).Find(&comments)
	return comments, d.Error
}