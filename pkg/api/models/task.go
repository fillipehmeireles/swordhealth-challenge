package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title        string    `gorm:"not null;size:255"`
	Summary      string    `gorm:"not null; size:2500"`
	Status       bool      `gorm:"default:false"` // 'True' when the task is performed
	PerformedAt  time.Time `gorm:"default:null"`
	Technician   Technician
	TechnicianID int
}
