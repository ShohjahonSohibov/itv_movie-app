package config

import (
	"itv_movie_app/internal/models"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() (*models.Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	expiryHour, _ := strconv.Atoi(os.Getenv("JWT_EXPIRY_HOUR"))

	return &models.Config{
		Server: models.ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: models.DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
		JWT: models.JWTConfig{
			Secret:     os.Getenv("JWT_SECRET"),
			ExpiryHour: expiryHour,
		},
	}, nil
}
