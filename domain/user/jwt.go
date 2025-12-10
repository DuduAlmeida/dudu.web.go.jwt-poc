package user

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenDuration  = time.Minute * 15   // 15 minutos
	refreshTokenDuration = time.Hour * 24 * 7 // 7 dias
)

type UserClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (user *User) GenerateAccessToken(jwtKeySecret string) (string, error) {
	expirationTime := time.Now().Add(accessTokenDuration)
	claims := &UserClaims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKeySecret))
}

func (user *User) GenerateRefreshToken(jwtKeySecret string) (string, error) {
	expirationTime := time.Now().Add(refreshTokenDuration)
	claims := &jwt.RegisteredClaims{
		Subject:   user.ID,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKeySecret))
}
