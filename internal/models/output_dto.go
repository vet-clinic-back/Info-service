package models

type ErrorDTO struct {
	Message string `json:"message"`
}

type OutputPetDTO struct {
	Pet     Pet  `json:"pet_info"`
	OwnerID uint `json:"owner_id"`
	VetID   uint `json:"vet_id"`
}
