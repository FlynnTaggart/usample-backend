package servers

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"user-service/internal/models"
	"user-service/internal/pb"
)

func (s UsersServer) AddUserLink(_ context.Context, req *pb.UserLink) (*pb.DefaultResponse, error) {
	userId, err := uuid.Parse(req.Id)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid user id",
		}, nil
	}

	link := &models.UserLink{
		Url:    req.Url,
		Type:   req.Type,
		UserId: userId,
	}

	err = s.service.AddUserLink(link)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}

	return &pb.DefaultResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s UsersServer) GetUserLinks(_ context.Context, req *pb.GetUserLinksRequest) (*pb.UserLinksResponse, error) {
	res, err := s.service.GetUserLinks()
	if err != nil {
		return &pb.UserLinksResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}

	resp := &pb.UserLinksResponse{
		Status:    http.StatusOK,
		UserLinks: []*pb.UserLink{},
	}

	for _, v := range res {
		resp.UserLinks = append(resp.UserLinks, &pb.UserLink{
			Id:     v.Id.String(),
			Type:   v.Type,
			Url:    v.Url,
			UserId: v.UserId.String(),
		})
	}

	return resp, nil
}

func (s UsersServer) DeleteUserLink(_ context.Context, req *pb.DeleteUserLinkRequest) (*pb.DefaultResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid user id",
		}, nil
	}
	linkId, err := uuid.Parse(req.Id)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid request id",
		}, nil
	}

	err = s.service.DeleteUserLink(userId, linkId)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}

	return &pb.DefaultResponse{
		Status: http.StatusOK,
	}, nil
}
