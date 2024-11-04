package http_utils

import (
	"github.com/gin-gonic/gin"
	"github.com/vet-clinic-back/info-service/internal/models"
	"strconv"
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

// getUint64Param returns *uint param. On error returns error and nil if param not exists
func getUint64Param(param string, c *gin.Context) (*uint, error) {
	stringParam, ok := c.GetQuery(param)
	if ok {
		paramUint64, err := strconv.ParseUint(stringParam, 10, 32)
		if err != nil {
			return nil, err
		}
		result := uint(paramUint64)
		return &result, nil
	} else {
		return nil, nil
	}
}
