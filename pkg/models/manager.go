package models

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	Email    string
	Password string
	Name     string
	//	CreatedAt time.Time `gorm:"autoCreateTime:true"`
	//	UpdatedAt time.Time `gorm:"autoUpdateTime:true"`
}
