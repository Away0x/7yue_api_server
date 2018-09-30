// 书籍短评
package model

import "time"

type BookComment struct {
	Id uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"pubdate"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`

	Content string `gorm:"varchar(128)" json:"content"`
	BookRefer uint
}

func (BookComment) TableName() string {
	return "book_comment"
}