package models

import (
	"github.com/google/uuid"
)

type SampleAccessType int32

const (
	ALL SampleAccessType = iota
	FRIENDS
	PRIVATE
)

type UserType int32

const (
	DEFAULT UserType = iota
	ADMIN
)

type LinkType int32

//goland:noinspection GoSnakeCaseUsage
const (
	CUSTOM_WEBSITE LinkType = iota
	SOUNDCLOUD
	VK
)

type User struct {
	Id                uuid.UUID
	Nickname          string
	FirstName         string
	SecondName        string
	DefaultAccessType SampleAccessType
	UserType          UserType
	Bio               string
}

type UserLink struct {
	Id     uuid.UUID
	Type   LinkType
	Url    string
	UserId uuid.UUID
}

type FriendRequest struct {
	Id         uuid.UUID
	SenderId   uuid.UUID
	ReceiverId uuid.UUID
	IsAccepted bool
}
