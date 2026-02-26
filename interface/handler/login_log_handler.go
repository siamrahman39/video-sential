package handler

import (
	"video-sentinel/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginLogHandler struct {
	svc *application.LoginLogService
}

func NewLoginLogHandler(svc *application.LoginLogService) *LoginLogHandler {
	return &LoginLogHandler{svc: svc}
}

func (h *LoginLogHandler) List(c *gin.Context) {
	logs, err := h.svc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "db error"}); return
	}
	c.JSON(http.StatusOK, gin.H{"data": logs})
}