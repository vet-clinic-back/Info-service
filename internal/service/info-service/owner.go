package infoservice

import "github.com/vet-clinic-back/info-service/internal/models"

func (s *InfoService) CreateOwner(owner models.Owner) (uint, error) {
	return s.storage.CreateOwner(owner)
}

func (s *InfoService) GetOwner(owner models.Owner) (models.Owner, error) {
	return s.storage.GetOwner(owner)
}

func (s *InfoService) GetAllOwners() ([]models.Owner, error) {
	return s.storage.GetAllOwners()
}

func (s *InfoService) UpdateOwner(owner models.Owner) (models.Owner, error) {
	return s.storage.UpdateOwner(owner)
}

func (s *InfoService) DeleteOwner(id uint) error {
	return s.storage.DeleteOwner(id)
}
