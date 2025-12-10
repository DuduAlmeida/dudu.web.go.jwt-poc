package controller

import (
	"net/http"

	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/api/common"
	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/api/middlewares"
	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/domain/user"
	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/mocks"
	"github.com/labstack/echo/v4"
)

type CarController struct {
	accessTokenSecret string
}

func NewCarController(accessTokenSecret string) Controller {
	return &CarController{
		accessTokenSecret,
	}
}
func (c *CarController) GetAllCars(ctx echo.Context) error {
	userClaims := common.GetEchoContextValue[user.UserClaims](ctx, common.UserClaimsKey)

	if userClaims == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro de processamento de autenticação")
	}

	return ctx.JSON(http.StatusOK, mocks.CarsMocked)
}

func (c *CarController) RegisterRoutes(g *echo.Group) {
	g.GET("/cars", c.GetAllCars, middlewares.JWTMiddleware(c.accessTokenSecret))
}
