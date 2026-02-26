package application

import (
	"video-sentinel/domain"
	"gorm.io/gorm"
)

type LoginLogService struct {
	db *gorm.DB
}

func NewLoginLogService(db *gorm.DB) *LoginLogService { return &LoginLogService{db: db} }

func (s *LoginLogService) List() ([]domain.LoginLog, error) {
	var logs []domain.LoginLog
	err := s.db.Preload("User").Order("created_at desc").Find(&logs).Error
	return logs, err
}