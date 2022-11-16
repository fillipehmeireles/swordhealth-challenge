package dtos

import "time"

type Message struct {
	TechName    string
	TechID      int
	TaskTitle   string
	TaskID      int
	PerformedAt time.Time
}
