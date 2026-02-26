package application

import (
	"video-sentinel/domain"
	"video-sentinel/infra/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService { return &UserService{db: db} }

func (s *UserService) Register(first, last, email, plainPass string, shopOwner bool) (*domain.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(plainPass), bcrypt.DefaultCost)
	u := domain.User{FirstName: first, LastName: last, Email: email, Password: string(hash), ShopOwner: shopOwner}
	if err := s.db.Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *UserService) Login(email, plainPass, ip string) (string, error) {
	var u domain.User
	if err := s.db.Where("email = ?", email).First(&u).Error; err != nil {
		return "", err
	}
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPass)) != nil {
		return "", bcrypt.ErrMismatchedHashAndPassword
	}
	s.db.Create(&domain.LoginLog{UserID: u.ID, IP: ip})
	return jwt.GenerateToken(u.ID)
}