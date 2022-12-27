package users

import (
	"api-gateway-service/internal/auth"
	"api-gateway-service/internal/samples/handlers"
	"api-gateway-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func InitializeSamplesRoutes(a fiber.Router, URL string, logger logger.Logger, authClient *auth.ServiceClient) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(URL, logger),
	}

	m := auth.InitAuthMiddleware(authClient, logger)

	samplePublicGroup := a.Group("/samples")
	samplePublicGroup.Get("/", svc.GetSamples)
	samplePublicGroup.Get("/prefix", svc.GetSamplesByNamePrefix)
	samplePublicGroup.Get("/:id", svc.GetSampleData)
	samplePublicGroup.Get("/:id/file", svc.GetSampleFile)
	//samplePublicGroup.Get("/:id/likes", svc.GetSampleLikes)
	//samplePublicGroup.Get("/:id/usages", svc.GetAllSampleUsages)
	//samplePublicGroup.Get("/:id/comments", svc.GetAllCommentsFromSample)
	//samplePublicGroup.Get("/:id/comments/:comment_id/likes", svc.GetCommentLikes)

	a.Get("/users/:id/samples", svc.GetUserSamples)

	sampleProtectedGroup := samplePublicGroup.Group("/protected", m.AuthRequired)
	sampleProtectedGroup.Get("/", svc.GetSamples)
	sampleProtectedGroup.Get("/prefix", svc.GetSamplesByNamePrefix)
	sampleProtectedGroup.Get("/:id", svc.GetSampleData)
	sampleProtectedGroup.Get("/:id/file", svc.GetSampleFile)
	sampleProtectedGroup.Post("/", svc.UploadSample)
	//sampleProtectedGroup.Get("/:id/likes", svc.GetSampleLikes)
	//sampleProtectedGroup.Post("/:id/likes", svc.ToggleSampleLike)
	//sampleProtectedGroup.Get("/:id/usages", svc.GetAllSampleUsages)
	//sampleProtectedGroup.Get("/:id/comments", svc.GetAllCommentsFromSample)
	//sampleProtectedGroup.Get("/:id/comments/:comment_id/likes", svc.GetCommentLikes)
	//
	//coverPublicGroup := a.Group("/covers")
	//coverPublicGroup.Get("/:id", svc.GetCover)
	//coverProtectedGroup := coverPublicGroup.Group("/protected", m.AuthRequired)
	//coverProtectedGroup.Delete("/:id", svc.DeleteCover)
	//coverProtectedGroup.Get("/:id", svc.GetCover)
	//coverProtectedGroup.Post("/", svc.UploadCover)
	//
	//sampleUsagesGroup := a.Group("/usages", m.AuthRequired)
	//sampleUsagesGroup.Post("/", svc.AddSampleUsage)
	//sampleUsagesGroup.Delete("/:id", svc.DeleteSampleUsage)
	//sampleUsagesGroup.Put("/:id", svc.EditSampleUsage)
	//
	//sampleCommentsGroup := a.Group("/comments", m.AuthRequired)
	//sampleCommentsGroup.Post("/", svc.AddCommentToSample)
	//sampleCommentsGroup.Delete("/:id", svc.DeleteComment)
	//sampleCommentsGroup.Put("/:id", svc.EditComment)
	//sampleCommentsGroup.Post("/:id/likes", svc.ToggleCommentLike)

	return svc
}

func (svc ServiceClient) GetSamples(ctx *fiber.Ctx) error {
	return handlers.GetSamples(ctx, svc.Client)
}

func (svc ServiceClient) GetSamplesByNamePrefix(ctx *fiber.Ctx) error {
	return handlers.GetSamplesByNamePrefix(ctx, svc.Client)
}

func (svc ServiceClient) GetSampleData(ctx *fiber.Ctx) error {
	return handlers.GetSampleData(ctx, svc.Client)
}

func (svc ServiceClient) GetSampleFile(ctx *fiber.Ctx) error {
	return handlers.GetSampleFile(ctx, svc.Client)
}

func (svc ServiceClient) UploadSample(ctx *fiber.Ctx) error {
	return handlers.UploadSample(ctx, svc.Client)
}

func (svc ServiceClient) GetUserSamples(ctx *fiber.Ctx) error {
	return handlers.GetUserSamples(ctx, svc.Client)
}
