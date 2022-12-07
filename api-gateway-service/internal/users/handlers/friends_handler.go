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

func GetUserFriends(ctx *fiber.Ctx, client users_pb.UsersServiceClient) error {
	userId := fmt.Sprintf("%v", ctx.Locals("userId"))

	res, err := client.GetUserFriends(context.Background(), &users_pb.GetUserFriendsRequest{UserId: userId})

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func GetUserSentFriends(ctx *fiber.Ctx, client users_pb.UsersServiceClient) error {
	userId := fmt.Sprintf("%v", ctx.Locals("userId"))

	res, err := client.GetUserSentFriends(context.Background(), &users_pb.GetUserFriendsRequest{UserId: userId})

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func GetUserReceivedFriends(ctx *fiber.Ctx, client users_pb.UsersServiceClient) error {
	userId := fmt.Sprintf("%v", ctx.Locals("userId"))

	res, err := client.GetUserReceivedFriends(context.Background(), &users_pb.GetUserFriendsRequest{UserId: userId})

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func SendFriend(ctx *fiber.Ctx, client users_pb.UsersServiceClient) error {
	body := users_pb.SendFriendRequest{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	if !validator.ValidateUUID(body.ReceiverId) {
		return utils.ReturnBadRequest(errors.New("gateway: send friend: invalid receiver id"), ctx, fiber.StatusBadRequest)
	}

	body.SenderId = fmt.Sprintf("%v", ctx.Locals("userId"))

	res, err := client.SendFriend(context.Background(), &body)

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.SendStatus(int(res.Status))
}

func AcceptFriend(ctx *fiber.Ctx, client users_pb.UsersServiceClient) error {
	body := users_pb.AcceptDeclineFriendRequest{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	if !validator.ValidateUUID(body.RequestId) {
		return utils.ReturnBadRequest(errors.New("gateway: accept friend: invalid request id"), ctx, fiber.StatusBadRequest)
	}

	body.UserId = fmt.Sprintf("%v", ctx.Locals("userId"))

	res, err := client.AcceptFriend(context.Background(), &body)

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.SendStatus(int(res.Status))
}

func DeclineFriend(ctx *fiber.Ctx, client users_pb.UsersServiceClient) error {
	body := users_pb.AcceptDeclineFriendRequest{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	if !validator.ValidateUUID(body.RequestId) {
		return utils.ReturnBadRequest(errors.New("gateway: decline friend: invalid request id"), ctx, fiber.StatusBadRequest)
	}

	body.UserId = fmt.Sprintf("%v", ctx.Locals("userId"))

	res, err := client.DeclineFriend(context.Background(), &body)

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	return ctx.SendStatus(int(res.Status))
}
