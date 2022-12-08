package validator

import "github.com/google/uuid"

func ValidateUUID(id string) bool {
	_, err := uuid.Parse(id)
	return nil == err
}
