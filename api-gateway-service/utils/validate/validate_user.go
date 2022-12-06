package validate

import (
	"api-gateway-service/internal/pb/users_pb"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

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
	return nil
}
