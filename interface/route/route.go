package route

import (
	"video-sentinel/application"
	"video-sentinel/interface/handler"
	"video-sentinel/infra/jwt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(r *gin.Engine, db *gorm.DB) {
	userSvc := application.NewUserService(db)
	loginLogSvc := application.NewLoginLogService(db)

	authH := handler.NewAuthHandler(userSvc)
	logH := handler.NewLoginLogHandler(loginLogSvc)

	// public
	r.POST("/register", authH.Register)
	r.POST("/login", authH.Login)

	// protected
	protected := r.Group("/api")
	protected.Use(JWTMiddleware())
	protected.GET("/login-list", logH.List)
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(401, gin.H{"msg": "missing token"}); c.Abort(); return
		}
		userID, err := jwt.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(401, gin.H{"msg": "invalid token"}); c.Abort(); return
		}
		c.Set("userID", userID)
		c.Next()
	}
}