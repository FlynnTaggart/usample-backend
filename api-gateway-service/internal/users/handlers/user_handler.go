package handlers

import (
	"api-gateway-service/internal/pb/users_pb"
	"api-gateway-service/utils"
	"api-gateway-service/utils/validate"
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx, client users_pb.UsersServiceClient) error {
	body := users_pb.User{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	if err := validate.ValidateUser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	res, err := client.CreateUser(context.Background(), &body)

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.SendStatus(int(res.Status))
}

func GetUsers(ctx *fiber.Ctx, client users_pb.UsersServiceClient) error {
	body := users_pb.GetUsersRequest{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	res, err := client.GetUsers(context.Background(), &body)

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func GetUser(ctx *fiber.Ctx, client users_pb.UsersServiceClient) error {
	id := ctx.Params("id")

	if len(id) == 0 || !validate.ValidateUUID(id) {
		return utils.ReturnBadRequest(errors.New("gateway: get user: invalid id"), ctx, fiber.StatusBadRequest)
	}

	res, err := client.GetUser(context.Background(), &users_pb.GetUserRequest{
		Id: id,
	})

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func GetUsersByNicknamePrefix(ctx *fiber.Ctx, client users_pb.UsersServiceClient) error {
	body := users_pb.GetUsersByNicknamePrefixRequest{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	res, err := client.GetUsersByNicknamePrefix(context.Background(), &body)

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func GetUserByNickname(ctx *fiber.Ctx, client users_pb.UsersServiceClient) error {
	body := users_pb.GetUserByNicknameRequest{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	res, err := client.GetUserByNickname(context.Background(), &body)

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func UpdateUserInfo(ctx *fiber.Ctx, client users_pb.UsersServiceClient) error {
	body := users_pb.User{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	if err := validate.ValidateUser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	res, err := client.CreateUser(context.Background(), &body)

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.SendStatus(int(res.Status))
}
