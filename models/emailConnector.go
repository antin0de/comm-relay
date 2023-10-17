package models

import (
	"gorm.io/gorm"
)

type EmailConnector struct {
	gorm.Model
	Name          string
	TargetAddress string
}
