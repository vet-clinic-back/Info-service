package http_utils

import (
	"github.com/gin-gonic/gin"
	"github.com/vet-clinic-back/info-service/internal/models"
)

func ParsePetFilters(c *gin.Context) (models.PetReqFilter, error) {
	var filters models.PetReqFilter

	petID, err := getUint64Param("pet_id", c)
	if err != nil {
		return filters, err
	}
	filters.PetID = petID

	vetID, err := getUint64Param("vet_id", c)
	if err != nil {
		return filters, err
	}
	filters.VetID = vetID

	ownerID, err := getUint64Param("owner_id", c)
	if err != nil {
		return filters, err
	}
	filters.OwnerID = ownerID

	offset, err := getUint64Param("offset", c)
	if err != nil {
		return filters, err
	}
	filters.Offset = offset

	limit, err := getUint64Param("limit", c)
	if err != nil {
		return filters, err
	}
	filters.Limit = limit

	return filters, nil
}
