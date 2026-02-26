package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DB        DBConfig
	HTTPPort  string
	JWTSecret string
}

type DBConfig struct {
	Host, Port, User, Pass, Name string
}

func Load() Config {
	godotenv.Load()
	return Config{
		DB: DBConfig{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASSWORD"),
			Name: os.Getenv("DB_NAME"),
		},
		HTTPPort:  os.Getenv("HTTP_PORT"),
		JWTSecret: os.Getenv("JWT_SECRET_KEY"),
	}
}