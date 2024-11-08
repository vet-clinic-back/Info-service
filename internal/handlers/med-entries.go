package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/vet-clinic-back/info-service/internal/models"
	http_utils "github.com/vet-clinic-back/info-service/internal/utils/http-utils"
	"net/http"
)

// @Summary Create med entry
// @Description Creates a new med entry
// @Security ApiKeyAuth
// @Tags MedEntry
// @Accept json
// @Produce json
// @Param input body models.MedicalEntry true "entry data"
// @Success 201 {object} number "Successfully created утекн"
// @Failure 400 {object} models.ErrorDTO "Invalid input body"
// @Failure 500 {object} models.ErrorDTO "Internal server error"
// @Router /info/v1/record/entries [post]
func (h *Handler) createEntry(c *gin.Context) {
	log := h.log.WithField("op", "Handler.createEntry")

	//petID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	//if err != nil {
	//	log.Error("invalid pet ID: ", err.Error())
	//	h.newErrorResponse(c, http.StatusBadRequest, "invalid pet ID")
	//	return
	//}

	var input models.MedicalEntry

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error("failed to parse json", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid json body")
		return
	}

	id, err := h.service.MedInfo.CreateMedEntry(input)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23503" {
				log.Errorf("foreign key constraint failed: %v", err)
				h.newErrorResponse(c, http.StatusBadRequest, "foreign key constraint failed. U use correct ids?")
				return
			}
		}
		log.Error("failed to create med entry", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, id) // ну ты сам начал тут не json возвращать, я продолжу)
}

// @Summary getEntries
// @Description Creates a new med entry
// @Security ApiKeyAuth
// @Tags MedEntry
// @Accept json
// @Produce json
// @Param entry_id query int false "Entry ID"
// @Param pet_id query int false "Pet ID"
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {object} []models.MedicalEntry "Successfully created утекн"
// @Failure 400 {object} models.ErrorDTO "failed to parse filters"
// @Failure 404 {object} models.ErrorDTO "Not found"
// @Failure 500 {object} models.ErrorDTO "Internal server error"
// @Router /info/v1/record/entries [get]
func (h *Handler) getEntries(c *gin.Context) {
	log := h.log.WithField("op", "Handler.getEntries")

	//petID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	//if err != nil {
	//	log.Error("invalid pet ID: ", err.Error())
	//	h.newErrorResponse(c, http.StatusBadRequest, "invalid pet ID")
	//	return
	//}

	filters, err := http_utils.ParseEntryFilters(c)
	if err != nil {
		log.Error("failed to parse filters: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "failed to parse filters")
		return
	}
	log.Debug("parsed filters", filters)

	entries, err := h.service.MedInfo.GetMedEntries(filters)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error("pets not found: ", err.Error())
			h.newErrorResponse(c, http.StatusNotFound, "not found")
			return
		}
		log.Error("failed to get entries", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, entries) // ну ты сам начал тут не json возвращать, я продолжу)
}
