package models

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string `json:"-"`
	Name     string
}
