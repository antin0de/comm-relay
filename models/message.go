package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ChannelID int
	Channel   Channel
	Title     string
	Content   string
}
