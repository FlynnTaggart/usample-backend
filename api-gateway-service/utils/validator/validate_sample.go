package validator

import (
	"api-gateway-service/internal/pb/samples_pb"
	"errors"
)

func ValidateSample(sample *samples_pb.SampleData) error {
	if len(sample.Name) == 0 {
		return errors.New("validate sample: empty sample name")
	}
	if len(sample.AuthorId) == 0 || !ValidateUUID(sample.AuthorId) {
		return errors.New("validate sample: invalid author id")
	}
	if len(sample.CoverId) > 0 && !ValidateUUID(sample.CoverId) {
		return errors.New("validate sample: invalid cover id")
	}
	return ValidateSampleAccessType(sample.AccessType)
}
