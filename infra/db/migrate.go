package db

import (
	"video-sentinel/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&domain.User{}, &domain.LoginLog{})
}