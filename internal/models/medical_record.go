package models

type MedicalRecord struct {
	vetID    uint `json:"vet_id"`
	ownderID uint `json:"owner_id"`
	petID    uint `json:"pet_id"`
}
