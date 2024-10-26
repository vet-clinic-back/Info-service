package infoservice

import "github.com/vet-clinic-back/info-service/internal/models"

func (s *InfoService) CreatePet(pet models.Pet) (uint, error) {
	return s.storage.CreatePet(pet)
}

func (s *InfoService) GetPet(pet models.Pet) (models.Pet, error) {
	return s.storage.GetPet(pet)
}

func (s *InfoService) GetAllPets() ([]models.Pet, error) {
	return s.storage.GetAllPets()
}

func (s *InfoService) UpdatePet(pet models.Pet) (models.Pet, error) {
	return s.storage.UpdatePet(pet)
}

func (s *InfoService) DeletePet(id uint) error {
	return s.storage.DeletePet(id)
}
