package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vet-clinic-back/info-service/internal/models"
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
		log.Error("failed to create med entry", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, id) // ну ты сам начал тут не json возвращать, я продолжу)
}
