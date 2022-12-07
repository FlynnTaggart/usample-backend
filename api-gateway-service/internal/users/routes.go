package users

import (
	"api-gateway-service/internal/auth"
	"api-gateway-service/internal/users/handlers"
	"api-gateway-service/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func InitializeUsersRoutes(a fiber.Router, URL string, logger logger.Logger, authClient *auth.ServiceClient) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(URL, logger),
	}

	m := auth.InitAuthMiddleware(authClient, logger)

	getUsersGroup := a.Group("/users")

	getUsersGroup.Get("/", svc.GetUsers)
	getUsersGroup.Get("/:id", svc.GetUser)
	getUsersGroup.Get("/prefix", svc.GetUsersByNicknamePrefix)
	getUsersGroup.Get("/nickname", svc.GetUserByNickname)
	getUsersGroup.Get("/:id/friends", svc.GetUserFriends)

	protectedUsersGroup := getUsersGroup.Group("/protected", m.AuthRequired)
	protectedUsersGroup.Post("/", svc.CreateUser)
	protectedUsersGroup.Patch("/", svc.UpdateUserInfo)

	userLinkGroup := protectedUsersGroup.Group("/links")
	userLinkGroup.Post("/", svc.AddUserLink)
	userLinkGroup.Get("/", svc.GetUserLinks)
	userLinkGroup.Delete("/", svc.DeleteUserLink)

	userFriendsGroup := protectedUsersGroup.Group("/friends")
	userFriendsGroup.Get("/sent", svc.GetUserSentFriends)
	userFriendsGroup.Get("/received", svc.GetUserReceivedFriends)
	userFriendsGroup.Post("/", svc.SendFriend)
	userFriendsGroup.Post("/accept", svc.AcceptFriend)
	userFriendsGroup.Post("/decline", svc.DeclineFriend)

	return svc
}

func (svc ServiceClient) GetUsers(ctx *fiber.Ctx) error {
	return handlers.GetUsers(ctx, svc.Client)
}

func (svc ServiceClient) GetUser(ctx *fiber.Ctx) error {
	return handlers.GetUser(ctx, svc.Client)
}

func (svc ServiceClient) GetUsersByNicknamePrefix(ctx *fiber.Ctx) error {
	return handlers.GetUsersByNicknamePrefix(ctx, svc.Client)
}

func (svc ServiceClient) GetUserByNickname(ctx *fiber.Ctx) error {
	return handlers.GetUserByNickname(ctx, svc.Client)
}

func (svc ServiceClient) CreateUser(ctx *fiber.Ctx) error {
	return handlers.CreateUser(ctx, svc.Client)
}

func (svc ServiceClient) UpdateUserInfo(ctx *fiber.Ctx) error {
	return handlers.UpdateUserInfo(ctx, svc.Client)
}

func (svc ServiceClient) AddUserLink(ctx *fiber.Ctx) error {
	return handlers.AddUserLink(ctx, svc.Client)
}

func (svc ServiceClient) GetUserLinks(ctx *fiber.Ctx) error {
	return handlers.GetUserLinks(ctx, svc.Client)
}

func (svc ServiceClient) DeleteUserLink(ctx *fiber.Ctx) error {
	return handlers.DeleteUserLink(ctx, svc.Client)
}

func (svc ServiceClient) GetUserFriends(ctx *fiber.Ctx) error {
	return handlers.GetUserFriends(ctx, svc.Client)
}

func (svc ServiceClient) GetUserSentFriends(ctx *fiber.Ctx) error {
	return handlers.GetUserSentFriends(ctx, svc.Client)
}

func (svc ServiceClient) GetUserReceivedFriends(ctx *fiber.Ctx) error {
	return handlers.GetUserReceivedFriends(ctx, svc.Client)
}

func (svc ServiceClient) SendFriend(ctx *fiber.Ctx) error {
	return handlers.SendFriend(ctx, svc.Client)
}

func (svc ServiceClient) AcceptFriend(ctx *fiber.Ctx) error {
	return handlers.AcceptFriend(ctx, svc.Client)
}

func (svc ServiceClient) DeclineFriend(ctx *fiber.Ctx) error {
	return handlers.DeclineFriend(ctx, svc.Client)
}
