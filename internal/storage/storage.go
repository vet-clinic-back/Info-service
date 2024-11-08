package storage

import (
	"github.com/vet-clinic-back/info-service/internal/config"
	"github.com/vet-clinic-back/info-service/internal/logging"
	"github.com/vet-clinic-back/info-service/internal/models"
	"github.com/vet-clinic-back/info-service/internal/storage/postgres"
)

// Iterface to interact with user data
type Pet interface {
	CreatePetWithCard(pet models.Pet, ownderID uint, vetID uint) (uint, error)
	GetPet(pet models.Pet) (models.Pet, error)
	GetPetsWithOwnerAndVet(filter models.PetReqFilter) ([]models.OutputPetDTO, error)
	UpdatePet(pet models.Pet) (models.Pet, error)
	DelPetWithCard(id uint) error
}

type Owner interface {
	CreateOwner(user models.Owner) (uint, error)
	GetOwner(owner models.Owner) (models.Owner, error)
	GetAllOwners() ([]models.Owner, error)
	UpdateOwner(owner models.Owner) (models.Owner, error)
	DeleteOwner(id uint) error
}

type MedEntry interface {
	CreateMedEntry(entry models.MedicalEntry) (uint, error)
	DeleteMedEntry(medRecordID uint, entryID uint) error
	GetMedEntries(models.EntryReqFilter) ([]models.MedicalEntry, error)
}

type Info interface {
	Owner
	Pet
	MedEntry
}

type StorageProcess interface {
	Shutdown() error
}

type Storage struct {
	Info
	StorageProcess
}

func New(log *logging.Logger, cfg *config.DbConfig) *Storage {
	pg := postgres.New(log, cfg)
	return &Storage{
		Info:           pg,
		StorageProcess: pg,
	}
}
