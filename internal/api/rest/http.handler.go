package rest

import (
	"go-ecommerce-app/internal/helper"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestHandler struct {
	App  fiber.Router
	DB   *gorm.DB
	Auth helper.Auth
}
