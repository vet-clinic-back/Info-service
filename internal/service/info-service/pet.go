package infoservice

import "github.com/vet-clinic-back/info-service/internal/models"

func (s *InfoService) CreatePetWithCard(pet models.Pet, ownderID uint, vetID uint) (uint, error) {
	return s.storage.CreatePetWithCard(pet, ownderID, vetID)
}

func (s *InfoService) GetPet(pet models.Pet) (models.Pet, error) {
	return s.storage.GetPet(pet)
}

func (s *InfoService) GetPets(filter models.PetReqFilter) ([]models.OutputPetDTO, error) {
	return s.storage.GetPetsWithOwnerAndVet(filter)
}

func (s *InfoService) UpdatePet(pet models.Pet) (models.Pet, error) {
	return s.storage.UpdatePet(pet)
}

func (s *InfoService) DelPetWithCard(id uint) error {
	return s.storage.DelPetWithCard(id)
}
