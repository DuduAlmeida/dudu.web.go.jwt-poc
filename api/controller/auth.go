package controller

import (
	"log"
	"net/http"

	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/api/request"
	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/api/response"
	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/app/services"
	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/mocks"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	jwtService services.JwtService
}

func NewAuthController(jwtService services.JwtService) Controller {
	return &AuthController{
		jwtService,
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

	acessToken, err := c.jwtService.GenerateAccessToken(user.ID)
	if err != nil {
		log.Printf("Erro ao gerar tokens: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro interno ao gerar tokens")
	}
	refreshToken, err := c.jwtService.GenerateRefreshToken(user.ID)
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

	credentials, err := c.jwtService.GetRefreshCredentialsByToken(req.RefreshToken)

	if err != nil || credentials.Token == nil || !credentials.Token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Refresh token inválido ou expirado")
	}

	claims := credentials.RegisteredClaims
	if claims == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Claims do token inválidas")
	}
	userID := claims.Subject

	user := mocks.UsersMocked.FindByUserid(userID)
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Usuário inválido")
	}

	acessToken, err := c.jwtService.GenerateAccessToken(user.ID)
	if err != nil {
		log.Printf("Erro ao gerar novo token: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro interno ao gerar tokens")
	}
	refreshToken, err := c.jwtService.GenerateRefreshToken(user.ID)
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
