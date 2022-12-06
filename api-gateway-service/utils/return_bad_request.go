package utils

import "github.com/gofiber/fiber/v2"

func ReturnBadRequest(err error, ctx *fiber.Ctx, status int) error {
	return ctx.Status(status).JSON(fiber.Map{
		"message": err.Error(),
	})
}
