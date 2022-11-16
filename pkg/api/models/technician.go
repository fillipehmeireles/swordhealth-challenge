package models

import "gorm.io/gorm"

type Technician struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string `json:"-"`
	Name     string
}
