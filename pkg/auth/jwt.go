package auth

import (
	"itv_movie_app/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTManager struct {
	config models.JWTConfig
}

func NewJWTManager(config models.JWTConfig) *JWTManager {
	return &JWTManager{config: config}
}

func (m *JWTManager) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * time.Duration(m.config.ExpiryHour)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.config.Secret))
}

// func (m *JWTManager) ValidateToken(tokenString string) (uint, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("unexpected signing method")
// 		}
// 		return []byte(m.config.Secret), nil
// 	})

// 	if err != nil {
// 		return 0, err
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		return 0, errors.New("invalid token")
// 	}

// 	userID := uint(claims["user_id"].(float64))
// 	return userID, nil
// }
