package handlers

import (
	"api-gateway-service/internal/pb/auth_pb"
	"api-gateway-service/utils"

	"context"

	"github.com/gofiber/fiber/v2"
)

type EmailPasswordRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *fiber.Ctx, client auth_pb.AuthServiceClient) error {
	body := EmailPasswordRequestBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	res, err := client.Register(context.Background(), &auth_pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func Login(ctx *fiber.Ctx, client auth_pb.AuthServiceClient) error {
	body := EmailPasswordRequestBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	res, err := client.Login(context.Background(), &auth_pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}
