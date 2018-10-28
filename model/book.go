package model

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	BookId  int    `gorm:"type:integer;not null;column:book_id"`
	Author  string `gorm:"type:varchar(128);not null;column:author"`
	Image   string `gorm:"type:varchar(128);not null;column:image"`
	Title   string `gorm:"type:varchar(128);not null;column:title"`
	Isbn    string `gorm:"type:varchar(128);not null;column:isbn;index:isbn"`
}

func (Book) TableName() string {
	return "book"
}

func (b *Book) Create() error {
	return DB.Create(&b).Error
}

func GetAllBook() ([]*Book, error) {
	bs := make([]*Book, 0)
	d := DB.Find(&bs)
	return bs, d.Error
}

type BookSerializer struct {
	Author     string `json:"author"`
	FavNums    int    `json:"fav_nums"`
	Id         int    `json:"id"`
	Image      string `json:"image"`
	LikeStatus int    `json:"like_status"`
	Title      string `json:"title"`
	Isbn       string `json:"isbn"`
}

func (b *Book) Serializer(like_status int, fav_nums int) BookSerializer {
	s := BookSerializer{
		Author: b.Author,
		FavNums: fav_nums,
		Id: b.BookId,
		Image: b.Image,
		LikeStatus: like_status,
		Title: b.Title,
		Isbn: b.Isbn,
	}

	return s
}
