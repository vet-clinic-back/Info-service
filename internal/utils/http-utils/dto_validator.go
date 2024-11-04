package http_utils

import (
	"errors"

	"github.com/vet-clinic-back/info-service/internal/models"
)

var ErrInvalidInputBody = errors.New("invalid input body")

func ValidateCreatingPetDTO(dto models.Pet) error {
	if dto.AnimalType == "" || dto.Name == "" || dto.Gender == "" || dto.Age == 0 ||
		dto.Weight == 0 || dto.Condition == "" || dto.Behavior == "" ||
		dto.ResearchStatus == "" {
		return ErrInvalidInputBody
	}
	return nil
}

func ValidateCreatingOwnerDTO(dto models.Owner) error {
	if dto.FullName == "" || dto.Email == "" || dto.Phone == "" || dto.PasswordHash == "" {
		return ErrInvalidInputBody
	}
	return nil
}
