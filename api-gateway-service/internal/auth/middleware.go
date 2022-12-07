package auth

import (
	"api-gateway-service/internal/pb/auth_pb"

	"context"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddlewareConfig struct {
	svc    *ServiceClient
	logger *log.Logger
}

func InitAuthMiddleware(svc *ServiceClient, logger *log.Logger) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc: svc, logger: logger}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *fiber.Ctx) {
	authorization := ctx.GetReqHeaders()["Authorization"]

	if authorization == "" {
		err := ctx.SendStatus(fiber.StatusUnauthorized)
		if err != nil {
			c.logger.Println(err.Error())
		}
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		err := ctx.SendStatus(fiber.StatusUnauthorized)
		if err != nil {
			c.logger.Println(err.Error())
		}
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &auth_pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != fiber.StatusOK {
		err = ctx.SendStatus(fiber.StatusUnauthorized)
		if err != nil {
			c.logger.Println(err.Error())
		}
		return
	}

	ctx.Locals("userId", res.UserId)

	err = ctx.Next()
	if err != nil {
		c.logger.Println(err.Error())
		err = ctx.SendStatus(fiber.StatusUnauthorized)
		if err != nil {
			c.logger.Println(err.Error())
		}
		return
	}
}
