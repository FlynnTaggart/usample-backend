package handlers

import (
	"api-gateway-service/internal/pb/samples_pb"
	"api-gateway-service/utils"
	"api-gateway-service/utils/validator"
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func GetSamples(ctx *fiber.Ctx, client samples_pb.SamplesServiceClient) error {
	body := samples_pb.GetSamplesRequest{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	res, err := client.GetSamples(context.Background(), &body)

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusInternalServerError)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func GetSamplesByNamePrefix(ctx *fiber.Ctx, client samples_pb.SamplesServiceClient) error {
	body := samples_pb.GetSamplesByNamePrefixRequest{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	res, err := client.GetSamplesByNamePrefix(context.Background(), &body)

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusInternalServerError)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func GetSampleData(ctx *fiber.Ctx, client samples_pb.SamplesServiceClient) error {
	id := ctx.Params("id")

	if len(id) == 0 || !validator.ValidateUUID(id) {
		return utils.ReturnBadRequest(errors.New("gateway: get sample: invalid id"), ctx, fiber.StatusBadRequest)
	}

	res, err := client.GetSampleData(context.Background(), &samples_pb.GetSampleRequest{
		Id: id,
	})

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusInternalServerError)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func GetSampleFile(ctx *fiber.Ctx, client samples_pb.SamplesServiceClient) error {
	return nil
}

func UploadSample(ctx *fiber.Ctx, client samples_pb.SamplesServiceClient) error {
	fileHeader, err := ctx.FormFile("sample")
	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}
	file, err := fileHeader.Open()
	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

}

func GetUserSamples(ctx *fiber.Ctx, client samples_pb.SamplesServiceClient) error {
	return nil
}
