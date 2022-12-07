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
	getUsersGroup.Group("/create").Use(m.AuthRequired)

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
