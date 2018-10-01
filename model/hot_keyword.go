// 热搜关键字
package model

import "github.com/jinzhu/gorm"

type HotKeyword struct {
	gorm.Model
	Text string `gorm:"column:text;unique;index:text;not null"`
}

func (HotKeyword) TableName() string {
	return "hot_keyword"
}

func (h *HotKeyword) Create() error {
	return DB.Create(&h).Error
}

// 获取所有热搜关键字
func GetAllHotKeyWord() ([]string, error){
	h := make([]HotKeyword, 0)
	d := DB.Find(&h)
	if d.Error != nil {
		return nil, d.Error
	}

	s := make([]string, len(h))
	for index, item := range h {
		s[index] = item.Text
	}
	return s, d.Error
}