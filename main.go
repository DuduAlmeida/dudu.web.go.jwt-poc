package main

import (
	"fmt"
	"log"

	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/api/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// --- Variáveis de Configuração Global -> todo: transformar em env ---
const (
	accessTokenSecret  = "super-secreta-access-key-123-echo"
	refreshTokenSecret = "ultra-secreta-refresh-key-456-echo"
	apiPrefix          = "/api/v1"
)

func main() {
	e := echo.New()

	// Middleware padrão do Echo
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	authController := controller.NewAuthController(refreshTokenSecret, accessTokenSecret)
	carController := controller.NewCarController(accessTokenSecret)

	authController.RegisterRoutes(e.Group(apiPrefix))
	carController.RegisterRoutes(e.Group(apiPrefix))

	port := ":8080"
	fmt.Printf("Servidor Echo rodando em http://localhost%s\n", port)
	log.Fatal(e.Start(port))
}
