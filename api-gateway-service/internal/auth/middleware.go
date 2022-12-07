package auth

import (
	"api-gateway-service/internal/pb/auth_pb"
	"api-gateway-service/pkg/logger"

	"context"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddlewareConfig struct {
	svc    *ServiceClient
	logger logger.Logger
}

func InitAuthMiddleware(svc *ServiceClient, logger logger.Logger) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc: svc, logger: logger}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *fiber.Ctx) error {
	authorization := ctx.GetReqHeaders()["Authorization"]

	if authorization == "" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	res, err := c.svc.Client.Validate(context.Background(), &auth_pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != fiber.StatusOK {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	ctx.Locals("userId", res.UserId)

	err = ctx.Next()
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	return nil
}
