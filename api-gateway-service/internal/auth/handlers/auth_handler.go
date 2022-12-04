package handlers

import (
	"api-gateway-service/internal/pb/auth_pb"

	"context"

	"github.com/gofiber/fiber/v2"
)

type EmailPasswordRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func returnBadRequest(err error, ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": err.Error(),
	})
}

func Register(ctx *fiber.Ctx, client auth_pb.AuthServiceClient) error {
	body := EmailPasswordRequestBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return returnBadRequest(err, ctx)
	}

	res, err := client.Register(context.Background(), &auth_pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		return returnBadRequest(err, ctx)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func Login(ctx *fiber.Ctx, client auth_pb.AuthServiceClient) error {
	body := EmailPasswordRequestBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return returnBadRequest(err, ctx)
	}

	res, err := client.Login(context.Background(), &auth_pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		return returnBadRequest(err, ctx)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}
