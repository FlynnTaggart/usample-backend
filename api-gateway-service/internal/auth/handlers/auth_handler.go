package handlers

import (
	"api-gateway-service/internal/pb/auth_pb"
	"api-gateway-service/internal/pb/samples_pb"
	"api-gateway-service/internal/pb/users_pb"
	"api-gateway-service/utils"
	"api-gateway-service/utils/validator"
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type EmailPasswordRequestBody struct {
	Email             string                      `json:"email"`
	Password          string                      `json:"password"`
	Nickname          string                      `json:"nickname"`
	FirstName         string                      `json:"first_name"`
	SecondName        string                      `json:"second_name"`
	DefaultAccessType samples_pb.SampleAccessType `json:"default_access_type"`
	UserType          users_pb.UserType           `json:"user_type"`
	Bio               string                      `json:"bio"`
}

func Register(ctx *fiber.Ctx, authClient auth_pb.AuthServiceClient, usersClient users_pb.UsersServiceClient) error {
	body := EmailPasswordRequestBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	user := &users_pb.User{
		Nickname:          body.Nickname,
		FirstName:         body.FirstName,
		SecondName:        body.SecondName,
		DefaultAccessType: body.DefaultAccessType,
		UserType:          body.UserType,
		Bio:               body.Bio,
	}

	if err := validator.ValidateUser(user); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	reqCtx := context.Background()

	authRes, err := authClient.Register(reqCtx, &auth_pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadGateway)
	}

	if int(authRes.Status) != fiber.StatusCreated {
		return utils.ReturnBadRequest(errors.New(authRes.Error), ctx, int(authRes.Status))
	}

	user.Id = authRes.UserId

	createRes, err := usersClient.CreateUser(reqCtx, user)
	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadGateway)
	}

	return ctx.SendStatus(int(createRes.Status))
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
