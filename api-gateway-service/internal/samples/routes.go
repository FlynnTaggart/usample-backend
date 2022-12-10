package users

import (
	"api-gateway-service/internal/auth"
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
	samplePublicGroup.Post("/likes", svc.GetSampleLikes)

	sampleProtectedGroup := samplePublicGroup.Group("/protected", m.AuthRequired)
	sampleProtectedGroup.Get("/", svc.GetSamples)
	sampleProtectedGroup.Get("/prefix", svc.GetSamplesByNamePrefix)
	sampleProtectedGroup.Get("/:id", svc.GetSampleData)
	sampleProtectedGroup.Get("/:id/file", svc.GetSampleFile)
	sampleProtectedGroup.Post("/", svc.UploadSample)
	sampleProtectedGroup.Post("/likes", svc.GetSampleLikes)
	sampleProtectedGroup.Post("/likes", svc.ToggleSampleLike)

	coverPublicGroup := a.Group("/covers")
	coverPublicGroup.Get("/", svc.GetCover)
	coverProtectedGroup := a.Group("/protected", m.AuthRequired)
	coverProtectedGroup.Delete("/:id", svc.DeleteCover)
	coverProtectedGroup.Post("/", svc.UploadCover)

	sampleUsagesGroup := samplePublicGroup.Group("/usages", m.AuthRequired)
	sampleUsagesGroup.Post("/", svc.AddSampleUsage)
	sampleUsagesGroup.Delete("/:id", svc.DeleteSampleUsage)
	sampleUsagesGroup.Get("/", svc.GetAllSampleUsages)
	sampleUsagesGroup.Put("/:id", svc.EditSampleUsage)

	return svc
}
