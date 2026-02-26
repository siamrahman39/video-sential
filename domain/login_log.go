package domain

import "gorm.io/gorm"

type LoginLog struct {
	gorm.Model
	UserID uint
	IP     string
	User   User
}