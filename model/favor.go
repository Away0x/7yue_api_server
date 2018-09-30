package model

import "time"

type Favor struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`

	Type int             `gorm:"column:type;not null" json:"type"`
	TargetId uint64      `gorm:"column:target_id;not null" json:"target_id"`
}

func (Favor) TableName() string {
	return "favor"
}