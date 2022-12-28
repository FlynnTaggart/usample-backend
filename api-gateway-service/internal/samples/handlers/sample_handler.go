package handlers

import (
	"api-gateway-service/internal/pb/samples_pb"
	"api-gateway-service/utils"
	"api-gateway-service/utils/validator"
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"mime/multipart"
	"os"
)

func GetSamples(ctx *fiber.Ctx, client samples_pb.SamplesServiceClient) error {
	body := samples_pb.GetSamplesRequest{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	curUserId := fmt.Sprintf("%v", ctx.Locals("userId"))
	if len(curUserId) > 0 {
		body.UserId = curUserId
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

	curUserId := fmt.Sprintf("%v", ctx.Locals("userId"))
	if len(curUserId) > 0 {
		body.UserId = curUserId
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

	body := samples_pb.GetSampleRequest{
		Id: id,
	}

	curUserId := fmt.Sprintf("%v", ctx.Locals("userId"))
	if len(curUserId) > 0 {
		body.AuthorId = curUserId
	}

	res, err := client.GetSampleData(context.Background(), &body)

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusInternalServerError)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func GetSampleFile(ctx *fiber.Ctx, client samples_pb.SamplesServiceClient) error {
	id := ctx.Params("id")

	if len(id) == 0 || !validator.ValidateUUID(id) {
		return utils.ReturnBadRequest(errors.New("gateway: get sample: invalid id"), ctx, fiber.StatusBadRequest)
	}

	body := samples_pb.GetSampleRequest{
		Id: id,
	}

	curUserId := fmt.Sprintf("%v", ctx.Locals("userId"))
	if len(curUserId) > 0 {
		body.AuthorId = curUserId
	}

	stream, err := client.GetSampleFile(context.Background(), &body)
	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusInternalServerError)
	}

	res, err := stream.Recv()
	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusInternalServerError)
	}

	if len(res.SampleFile.GetFileType()) == 0 {
		return utils.ReturnBadRequest(errors.New("get sample file: empty response from internal service"), ctx, fiber.StatusInternalServerError)
	}

	fileName := fmt.Sprintf("./transfer/%s.%s", id, res.SampleFile.GetFileType())
	file, err := os.Create(fileName)
	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusInternalServerError)
	}
	writer := bufio.NewWriter(file)

	for {
		res, err = stream.Recv()
		if err == io.EOF {
			_ = writer.Flush()
			_ = file.Close()
			err = ctx.Status(int(res.Status)).SendFile(fileName)

			_ = os.Remove(fileName)

			return err
		}
		if err != nil {
			_ = writer.Flush()
			_ = file.Close()
			_ = os.Remove(fileName)
			return utils.ReturnBadRequest(err, ctx, fiber.StatusInternalServerError)
		}

		buf := res.GetSampleFile().GetContent()
		_, err = writer.Write(buf)
		if err != nil {
			_ = writer.Flush()
			_ = file.Close()
			_ = os.Remove(fileName)
			return utils.ReturnBadRequest(err, ctx, fiber.StatusInternalServerError)
		}
	}
}

func UploadSample(ctx *fiber.Ctx, client samples_pb.SamplesServiceClient) error {
	body := samples_pb.SampleData{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	if err := validator.ValidateSample(&body); err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}

	fileHeader, err := ctx.FormFile("sample")
	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}
	file, err := fileHeader.Open()
	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusBadRequest)
	}
	defer func(file multipart.File) {
		_ = file.Close()
	}(file)

	stream, err := client.UploadSample(context.Background())
	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusInternalServerError)
	}

	err = stream.Send(&samples_pb.UploadSampleRequest{
		Request: &samples_pb.UploadSampleRequest_SampleUploadData{
			SampleUploadData: &samples_pb.SampleUploadData{
				SampleData: &body,
			},
		},
	})
	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusInternalServerError)
	}

	buf := make([]byte, 1024) // maybe make env for chunk size

	for {
		num, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return utils.ReturnBadRequest(err, ctx, fiber.StatusInternalServerError)
		}

		if err := stream.Send(&samples_pb.UploadSampleRequest{
			Request: &samples_pb.UploadSampleRequest_Content{
				Content: buf[:num],
			},
		}); err != nil {
			return utils.ReturnBadRequest(err, ctx, fiber.StatusInternalServerError)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(int(res.Status))
}

func GetUserSamples(ctx *fiber.Ctx, client samples_pb.SamplesServiceClient) error {
	id := ctx.Params("id")

	if len(id) == 0 || !validator.ValidateUUID(id) {
		return utils.ReturnBadRequest(errors.New("gateway: get sample: invalid id"), ctx, fiber.StatusBadRequest)
	}

	body := samples_pb.GetUserSamplesRequest{
		AuthorId: id,
	}

	curUserId := fmt.Sprintf("%v", ctx.Locals("userId"))
	if len(curUserId) > 0 {
		body.UserId = curUserId
	}

	res, err := client.GetUserSamples(context.Background(), &body)

	if err != nil {
		return utils.ReturnBadRequest(err, ctx, fiber.StatusInternalServerError)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}
