package infoservice

import "github.com/vet-clinic-back/info-service/internal/models"

func (s *InfoService) CreateMedEntry(entry models.MedicalEntry) (uint, error) {
	return s.storage.CreateMedEntry(entry)
}
