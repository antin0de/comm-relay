package models

import (
	"gorm.io/gorm"
)

type EmailTarget struct {
	gorm.Model
	ChannelID    uint
	Channel      Channel
	EmailAddress string
}
