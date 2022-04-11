package model

import "gorm.io/gorm"

type WeChat struct {
	gorm.Model
	Content string `gorm:"type:varchar(50) not null"`
	Label   string `gorm:"type:varchar(10)"`
}
