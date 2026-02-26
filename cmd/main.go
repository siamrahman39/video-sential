package main

import (
	"video-sentinel/config"
	"video-sentinel/infra/db"
	"video-sentinel/interface/route"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	gormDB := db.MustConnect(cfg)
	db.Migrate(gormDB)

	r := gin.Default()
	route.Setup(r, gormDB)
	r.Run(":" + cfg.HTTPPort)
}