package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RespondWithError(ctx *fiber.Ctx, code int, err error) error {
	return ctx.Status(code).JSON(err.Error())
}

func RespondWithInternalError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
}

func RespondWithSucess(ctx *fiber.Ctx, code int, msg string, data any) error {
	return ctx.Status(code).JSON(&fiber.Map{
		"message": msg,
		"data":    data,
	})
}
