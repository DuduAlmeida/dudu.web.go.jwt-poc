package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenDuration  = time.Minute * 15   // 15 minutos
	refreshTokenDuration = time.Hour * 24 * 7 // 7 dias
)

type JwtService interface {
	GenerateAccessToken(subject string) (string, error)
	GenerateRefreshToken(subject string) (string, error)
	GetAccessCredentialsByToken(token string) (*JwtCredentials, error)
	GetRefreshCredentialsByToken(token string) (*JwtCredentials, error)
}

type jwtService struct {
	accessKeySecret  string
	refrestKeySecret string
}

func NewJwtService(
	accessKeySecret string,
	refrestKeySecret string,
) JwtService {
	return &jwtService{
		accessKeySecret:  accessKeySecret,
		refrestKeySecret: refrestKeySecret,
	}
}

func (s *jwtService) GenerateAccessToken(subject string) (string, error) {
	expirationTime := time.Now().Add(accessTokenDuration)
	claims := &jwt.RegisteredClaims{
		Subject:   subject,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.accessKeySecret))
}

func (s *jwtService) GenerateRefreshToken(subject string) (string, error) {
	expirationTime := time.Now().Add(refreshTokenDuration)
	claims := &jwt.RegisteredClaims{
		Subject:   subject,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.refrestKeySecret))
}

type JwtCredentials struct {
	*jwt.Token
	*jwt.RegisteredClaims
}

func (s *jwtService) GetAccessCredentialsByToken(token string) (*JwtCredentials, error) {
	return s.getCredentialsByToken(s.accessKeySecret, token)
}

func (s *jwtService) GetRefreshCredentialsByToken(token string) (*JwtCredentials, error) {
	return s.getCredentialsByToken(s.refrestKeySecret, token)
}

func (s *jwtService) getCredentialsByToken(keySecret string, receivedToken string) (*JwtCredentials, error) {
	claims := new(jwt.RegisteredClaims)
	token, err := jwt.ParseWithClaims(receivedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return []byte(keySecret), nil
	})

	if err != nil {
		return nil, err
	}

	return &JwtCredentials{
		Token:            token,
		RegisteredClaims: claims,
	}, nil
}
