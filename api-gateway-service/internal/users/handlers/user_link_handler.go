package handlers

import (
	"api-gateway-service/internal/pb/users_pb"
	"api-gateway-service/utils"
	"api-gateway-service/utils/validator"
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func AddUserLink(ctx *fiber.Ctx, client users_pb.UsersServiceClient) error {
	body := users_pb.UserLink{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	if err := validator.ValidateUserLink(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	body.UserId = fmt.Sprintf("%v", ctx.Locals("userId"))

	res, err := client.AddUserLink(context.Background(), &body)

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.SendStatus(int(res.Status))
}

func GetUserLinks(ctx *fiber.Ctx, client users_pb.UsersServiceClient) error {
	res, err := client.GetUserLinks(context.Background(), &users_pb.GetUserLinksRequest{})

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func DeleteUserLink(ctx *fiber.Ctx, client users_pb.UsersServiceClient) error {
	body := users_pb.DeleteUserLinkRequest{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	if !validator.ValidateUUID(body.Id) {
		return utils.ReturnBadRequest(errors.New("gateway: get user: invalid id"), ctx, fiber.StatusBadRequest)
	}

	body.UserId = fmt.Sprintf("%v", ctx.Locals("userId"))

	res, err := client.DeleteUserLink(context.Background(), &body)

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.SendStatus(int(res.Status))
}
