package api

import (
	"go-ecommerce-app/configs"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func StartServer(cfg configs.AppConfig) {
	port := cfg.ServerPort
	app := fiber.New()

	rh := &rest.RestHandler{
		App: app,
	}

	setupRoutes(rh)

	app.Listen(port)
}

func setupRoutes(restHand *rest.RestHandler) {
	controllers.UserRoutes(restHand)
}
