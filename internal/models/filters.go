package models

type PetReqFilter struct {
	// models.Pet add later
	PetID   *uint `json:"pet_id"`
	OwnerID *uint `json:"owner_id"`
	VetID   *uint `json:"vet_id"`
	Limit   *uint `json:"limit"`
	Offset  *uint `json:"offset"`
}

type EntryReqFilter struct {
	PetID   *uint `json:"pet_id"`
	EntryID *uint `json:"entry_id"`
	Limit   *uint `json:"limit"`
	Offset  *uint `json:"offset"`
}
