package model

import (
	"time"
)

type Classic struct {
	Id uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"pubdate"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`

	Content    string  `gorm:"type:varchar(255);not null;column:content" json:"content"`
	FavNums    int     `gorm:"type:integer;not null;column:fav_nums" json:"fav_nums"`
	Image      string  `gorm:"type:varchar(128);not null;column:image" json:"image"`
	Url        string  `gorm:"type:varchar(128);not null;column:url" json:"url"`
	Index      int     `gorm:"type:integer;not null;column:index" json:"index"`
	LikeStatus int     `gorm:"type:integer;not null;column:like_status" json:"like_status"`
	Title      string  `gorm:"type:varchar(128);not null;column:title" json:"title"`
	Type       int     `gorm:"type:integer;not null;column:type" json:"type"`
}

func (Classic) TableName() string {
	return "classic"
}
