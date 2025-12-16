package controller

import (
	"net/http"

	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/api/common"
	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/api/middlewares"
	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/app/services"
	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/mocks"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CarController struct {
	jwtService services.JwtService
}

func NewCarController(jwtService services.JwtService) Controller {
	return &CarController{
		jwtService,
	}
}
func (c *CarController) GetAllCars(ctx echo.Context) error {
	userClaims := common.GetEchoContextValue[jwt.RegisteredClaims](ctx, common.JwtRegisteredClaimsKey)

	if userClaims == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro de processamento de autenticação")
	}

	return ctx.JSON(http.StatusOK, mocks.CarsMocked)
}

func (c *CarController) RegisterRoutes(g *echo.Group) {
	g.GET("/cars", c.GetAllCars, middlewares.JWTMiddleware(c.jwtService))
}
