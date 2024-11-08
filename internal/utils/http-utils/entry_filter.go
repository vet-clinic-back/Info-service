package http_utils

import (
	"github.com/gin-gonic/gin"
	"github.com/vet-clinic-back/info-service/internal/models"
)

// @Param entry_id query int false "Entry ID"
// @Param pet_id query int false "Pet ID"

func ParseEntryFilters(c *gin.Context) (models.EntryReqFilter, error) {
	var filters models.EntryReqFilter

	petID, err := getUint64Param("pet_id", c)
	if err != nil {
		return filters, err
	}
	filters.PetID = petID

	entryID, err := getUint64Param("entry_id", c)
	if err != nil {
		return filters, err
	}
	filters.EntryID = entryID

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
