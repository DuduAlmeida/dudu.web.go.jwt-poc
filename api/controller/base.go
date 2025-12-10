package controller

import "github.com/labstack/echo/v4"

type Controller interface {
	RegisterRoutes(g *echo.Group)
}

//TODO: IMPLEMENT A BASE CONTROLLER WITH SOME RESPONSE UTILS
