package service

import (
	"github.com/vet-clinic-back/info-service/internal/logging"
	"github.com/vet-clinic-back/info-service/internal/models"
	infoservice "github.com/vet-clinic-back/info-service/internal/service/info-service"
	"github.com/vet-clinic-back/info-service/internal/storage"
)

type Info interface {
	CreatePet(pet models.Pet) (uint, error)
	GetPet(pet models.Pet) (models.Pet, error)
	GetAllPets() ([]models.Pet, error)
	UpdatePet(pet models.Pet) (models.Pet, error)
	DeletePet(id uint) error
	//
	CreateOwner(user models.Owner) (uint, error)
	GetOwner(owner models.Owner) (models.Owner, error)
	GetAllOwners() ([]models.Owner, error)
	UpdateOwner(owner models.Owner) (models.Owner, error)
	DeleteOwner(id uint) error
}

type Service struct {
	Info
}

func New(log *logging.Logger, stor storage.Info) *Service {
	return &Service{
		Info: infoservice.New(log, stor),
	}
}
