package infoservice

import "github.com/vet-clinic-back/info-service/internal/models"

func (s *InfoService) CreateMedEntry(entry models.MedicalEntry) (uint, error) {
	return s.storage.CreateMedEntry(entry)
}

func (s *InfoService) GetMedEntries(filters models.EntryReqFilter) ([]models.MedicalEntry, error) {
	return s.storage.GetMedEntries(filters)
}
