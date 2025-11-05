package api

import (
	"go-ecommerce-app/configs"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/routes"
	"go-ecommerce-app/internal/schema"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(cfg configs.AppConfig) {
	port := cfg.ServerPort
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})

	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	log.Printf("Database Connected")

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	db.AutoMigrate(&schema.User{})

	v1Routes := app.Group("/api/v1")

	v1Routes.Get("/health", healthCheck)

	rh := &rest.RestHandler{
		App: v1Routes,
		DB:  db,
	}

	setupRoutes(rh)
	app.Listen(port)
}

func healthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{
		"message": "Server is up and running",
	})
}

func setupRoutes(restHand *rest.RestHandler) {
	routes.UserRoutes(restHand)
}
