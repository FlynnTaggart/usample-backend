package models

import (
	"github.com/google/uuid"
	"user-service/internal/pb"
)

type User struct {
	Id                uuid.UUID
	Nickname          string
	FirstName         string
	SecondName        string
	DefaultAccessType pb.SampleAccessType
	UserType          pb.UserType
	Bio               string
}

type UserLink struct {
	Id     uuid.UUID
	Type   pb.LinkType
	Url    string
	UserId uuid.UUID
}

type FriendRequest struct {
	Id         uuid.UUID
	SenderId   uuid.UUID
	ReceiverId uuid.UUID
	IsAccepted bool
}
