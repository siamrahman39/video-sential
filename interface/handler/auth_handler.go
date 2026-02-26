package handler

import (
	"video-sentinel/application"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthHandler struct {
	userSvc *application.UserService
}

func NewAuthHandler(userSvc *application.UserService) *AuthHandler {
	return &AuthHandler{userSvc: userSvc}
}

func (h *AuthHandler) Register(c *gin.Context) {
	type Req struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		ShopOwner bool   `json:"shopowner"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
	}
	if _, err := h.userSvc.Register(req.FirstName, req.LastName, req.Email, req.Password, req.ShopOwner); err != nil {
		c.JSON(http.StatusConflict, gin.H{"msg": "email exists"}); return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "user created"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
	}
	token, err := h.userSvc.Login(req.Email, req.Password, c.ClientIP())
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "invalid credential"}); return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error"}); return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}