package models

import "gorm.io/gorm"

type Technician struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Name     string
}
