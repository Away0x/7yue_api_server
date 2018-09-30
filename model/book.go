package model

import "time"

/*
// 获取热门书籍 概要
{
		"author": "陈儒",
		"fav_nums": 0,
		"id": 18,
		"image": "https://img3.doubanio.com/lpic/s3435132.jpg",
		"like_status": 0,
		"title": "Python源码剖析"
}
 */

type Book struct {
	Id uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"pubdate"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`

	Comment []BookComment `gorm:"foreignkey:BookRefer"`
}

func (Book) TableName() string {
	return "book"
}
