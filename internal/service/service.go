package service

import (
	"github.com/vet-clinic-back/info-service/internal/logging"
	"github.com/vet-clinic-back/info-service/internal/models"
	infoservice "github.com/vet-clinic-back/info-service/internal/service/info-service"
	"github.com/vet-clinic-back/info-service/internal/storage"
)

type Info interface {
	CreatePetWithCard(pet models.Pet, ownderID uint, vetID uint) (uint, error)
	GetPet(pet models.Pet) (models.Pet, error)
	GetPets(filter models.PetReqFilter) ([]models.OutputPetDTO, error)
	UpdatePet(pet models.Pet) (models.Pet, error)
	DelPetWithCard(id uint) error
	// owner is used at auth service
	CreateOwner(user models.Owner) (uint, error)
	GetOwner(owner models.Owner) (models.Owner, error)
	GetAllOwners() ([]models.Owner, error)
	UpdateOwner(owner models.Owner) (models.Owner, error)
	DeleteOwner(id uint) error
}

type MedInfo interface {
	CreateMedEntry(entry models.MedicalEntry) (uint, error)
}

type Service struct {
	Info
	MedInfo
}

func New(log *logging.Logger, stor storage.Info) *Service {
	s := infoservice.New(log, stor)
	return &Service{
		Info:    s,
		MedInfo: s,
	}
}
