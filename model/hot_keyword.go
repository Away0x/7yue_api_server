// 热搜关键字
package model

import "time"

type HotKeyword struct {
	Id uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"pubdate"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`

	Text string `gorm:"column:text;unique;index:text;not null" json:"text"`
}