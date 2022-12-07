package validator

import (
	"api-gateway-service/internal/pb/samples_pb"
	"api-gateway-service/internal/pb/users_pb"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

func ValidateUserType(userType users_pb.UserType) error {
	switch userType {
	case users_pb.UserType_ADMIN, users_pb.UserType_DEFAULT:
		return nil
	}
	return errors.New("validate user: invalid user type")
}

func ValidateSampleAccessType(accessType samples_pb.SampleAccessType) error {
	switch accessType {
	case samples_pb.SampleAccessType_ALL, samples_pb.SampleAccessType_FRIENDS, samples_pb.SampleAccessType_PRIVATE:
		return nil
	}
	return errors.New("validate user: invalid default sample access type")
}

func ValidateUser(user *users_pb.User) error {
	if len(user.Id) == 0 {
		return errors.New("validate user: empty user id")
	}
	if len(user.Nickname) == 0 {
		return errors.New("validate user: empty user nickname")
	}
	_, err := uuid.Parse(user.Id)
	if err != nil {
		return fmt.Errorf("validate user: invalid user id: %v", err.Error())
	}
	err = ValidateUserType(user.UserType)
	if err != nil {
		return err
	}
	return ValidateSampleAccessType(user.DefaultAccessType)
}
