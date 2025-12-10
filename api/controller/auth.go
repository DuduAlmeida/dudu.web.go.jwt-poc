package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/api/request"
	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/api/response"
	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/mocks"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	refreshTokenSecret string
	accessTokenSecret  string
}

func NewAuthController(refreshTokenSecret string, accessTokenSecret string) Controller {
	return &AuthController{
		refreshTokenSecret,
		accessTokenSecret,
	}
}

func (c *AuthController) Login(ctx echo.Context) error {
	creds := new(request.LoginCredentials)
	if err := ctx.Bind(creds); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Requisição inválida")
	}

	user := mocks.UsersMocked.FindByUsername(creds.Username)
	if user == nil || user.Password != creds.Password {
		return echo.NewHTTPError(http.StatusUnauthorized, "Credenciais inválidas")
	}

	acessToken, err := user.GenerateAccessToken(c.accessTokenSecret)
	if err != nil {
		log.Printf("Erro ao gerar tokens: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro interno ao gerar tokens")
	}
	refreshToken, err := user.GenerateRefreshToken(c.refreshTokenSecret)
	if err != nil {
		log.Printf("Erro ao gerar tokens: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro interno ao gerar tokens")
	}

	return ctx.JSON(http.StatusOK, response.TokenPair{
		AccessToken:  acessToken,
		RefreshToken: refreshToken,
	})
}

func (c *AuthController) RefreshTokenHandler(ctx echo.Context) error {
	var req request.RefreshTokenCredentials

	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Requisição inválida")
	}

	token, err := jwt.ParseWithClaims(req.RefreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return []byte(c.refreshTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Refresh token inválido ou expirado")
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Claims do token inválidas")
	}
	userID := claims.Subject

	user := mocks.UsersMocked.FindByUserid(userID)
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Usuário inválido")
	}

	acessToken, err := user.GenerateAccessToken(c.accessTokenSecret)
	if err != nil {
		log.Printf("Erro ao gerar novo token: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro interno ao gerar tokens")
	}
	refreshToken, err := user.GenerateRefreshToken(c.refreshTokenSecret)
	if err != nil {
		log.Printf("Erro ao gerar novo token: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro interno ao gerar tokens")
	}

	return ctx.JSON(http.StatusOK, response.TokenPair{
		AccessToken:  acessToken,
		RefreshToken: refreshToken,
	})
}

func (c *AuthController) RegisterRoutes(g *echo.Group) {
	g.POST("/auth/login", c.Login)
	g.POST("/auth/refresh", c.RefreshTokenHandler)
}
