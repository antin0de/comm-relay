package models

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Channel{})
	db.AutoMigrate(&EmailConnector{})
	db.AutoMigrate(&Message{})
}
