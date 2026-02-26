package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secret = []byte("10202siam@") 
func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(secret)
}

func ValidateToken(signed string) (uint, error) {
	token, err := jwt.Parse(signed, func(t *jwt.Token) (interface{}, error) { return secret, nil })
	if err != nil || !token.Valid {
		return 0, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return uint(claims["sub"].(float64)), nil
}