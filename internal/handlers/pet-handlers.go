package handlers

import (
	"database/sql"
	"errors"
	"github.com/vet-clinic-back/info-service/internal/utils/http-utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vet-clinic-back/info-service/internal/models"
)

type createPetDTO struct {
	VetID   uint `json:"vet_id"`
	OwnerID uint `json:"owner_id"`
	models.Pet
}

// @Summary Create Pet
// @Description Create a new pet in the system. Age & weight should be > 0 & Gender should be 'Male' or 'Female'
// @Security ApiKeyAuth
// @Tags pets
// @Accept json
// @Produce json
// @Param input body createPetDTO true "Pet details"
// @Success 201 {object} number "Successfully created pet"
// @Failure 400 {object} models.ErrorDTO "Invalid input body"
// @Failure 500 {object} models.ErrorDTO "Internal server error"
// @Router /info/v1/pets [post]
func (h *Handler) createPet(c *gin.Context) {
	op := "Handler.createPet"
	log := h.log.WithField("op", op)

	var input createPetDTO

	log.Debug("binding json")
	if err := c.BindJSON(&input); err != nil {
		log.Error("failed to bind json: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	log.Debug("validating input")
	if err := http_utils.ValidateCreatingPetDTO(input.Pet); err != nil {
		log.Error("failed to validate input: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body. Age & weight should be > 0 & Gender "+
			"should be 'Male' or 'Female'")
		return
	}

	log.Debug("creating pet")
	pet, err := h.service.Info.CreatePetWithCard(input.Pet, input.OwnerID, input.VetID)
	if err != nil {
		log.Errorf("failed to create pet: %s", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to create pet")
		return
	}

	log.Info("successfully created pet")
	c.JSON(http.StatusCreated, pet)
}

// @Summary Get Pet
// @Description Get pet details by ID
// @Security ApiKeyAuth
// @Tags pets
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {object} models.Pet "Successfully retrieved pet"
// @Failure 404 {object} models.ErrorDTO "Pet not found"
// @Failure 500 {object} models.ErrorDTO "Internal server error"
// @Router /info/v1/pets/{id} [get]
func (h *Handler) getPet(c *gin.Context) {
	op := "Handler.getPet"
	log := h.log.WithField("op", op)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Error("invalid pet ID: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid pet ID")
		return
	}

	pt := models.Pet{ID: uint(id)}

	pet, err := h.service.Info.GetPet(pt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error("pet not found: ", err.Error())
			h.newErrorResponse(c, http.StatusNotFound, "pet not found")
			return
		}
		log.Error("failed to get pet: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to get pet")
		return
	}

	log.Info("successfully retrieved pet")
	c.JSON(http.StatusOK, pet)
}

// @Summary Get all pets
// @Description Get all pets details
// @Security ApiKeyAuth
// @Tags pets
// @Param pet_id query int false "Pet ID"
// @Param vet_id query int false "Veterinarian ID"
// @Param owner_id query int false "Owner ID"
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Produce json
// @Success 200 {object} models.OutputPetDTO "Successfully retrieved pets"
// @Failure 500 {object} models.ErrorDTO "Internal server error"
// @Router  /info/v1/pets [get]
func (h *Handler) getPets(c *gin.Context) {
	op := "Handler.getPets"
	log := h.log.WithField("op", op)

	filters, err := http_utils.ParsePetFilters(c)
	if err != nil {
		log.Error("failed to parse filters: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "failed to parse filters")
		return
	}

	log.WithField("filters", filters).Info("filters updated")

	log.Debug("retrieving all petsWithExtraInfo")
	petsWithExtraInfo, err := h.service.Info.GetPets(filters)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error("pets not found: ", err.Error())
			h.newErrorResponse(c, http.StatusNotFound, "not found")
			return
		}
		log.Error("failed to get petsWithExtraInfo with filter: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to get petsWithExtraInfo with filter")
		return
	}

	log.Info("successfully retrieved all petsWithExtraInfo")
	c.JSON(http.StatusOK, petsWithExtraInfo)
}

// @Summary Update Pet
// @Description Update pet details by ID
// @Security ApiKeyAuth
// @Tags pets
// @Accept json
// @Produce json
// @Param id path int true "Pet ID"
// @Param input body models.Pet true "Pet details"
// @Success 200 {object} models.Pet "Successfully updated pet"
// @Failure 400 {object} models.ErrorDTO "Invalid input body or pet ID"
// @Failure 404 {object} models.ErrorDTO "Pet not found"
// @Failure 500 {object} models.ErrorDTO "Internal server error"
// @Router /info/v1/pets/{id} [put]
func (h *Handler) updatePet(c *gin.Context) {
	op := "Handler.updatePet"
	log := h.log.WithField("op", op)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Error("invalid pet ID: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid pet ID")
		return
	}

	var input models.Pet
	if err := c.BindJSON(&input); err != nil {
		log.Error("failed to bind json: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	input.ID = uint(id)

	log.Debug("updating pet")
	updatedPet, err := h.service.Info.UpdatePet(input)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error("pet not found: ", err.Error())
			h.newErrorResponse(c, http.StatusNotFound, "pet not found")
			return
		}
		log.Error("failed to update pet: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to update pet")
		return
	}

	log.Info("successfully updated pet")
	c.JSON(http.StatusOK, updatedPet)
}

// @Summary Delete Pet
// @Description Delete pet details by ID
// @Security ApiKeyAuth
// @Tags pets
// @Accept json
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {object} models.Pet "Successfully deleted pet"
// @Failure 404 {object} models.ErrorDTO "Pet not found"
// @Failure 500 {object} models.ErrorDTO "Internal server error"
// @Router /info/v1/pets/{id} [delete]
func (h *Handler) deletePet(c *gin.Context) {
	op := "Handler.deletePet"
	log := h.log.WithField("op", op)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Error("invalid pet ID: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid pet ID")
		return
	}

	log.Debug("deleting pet")
	err = h.service.Info.DelPetWithCard(uint(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error("pet not found: ", err.Error())
			h.newErrorResponse(c, http.StatusNotFound, "pet not found")
			return
		}
		log.Error("failed to delete pet: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to delete pet")
		return
	}

	log.Info("successfully deleted pet")
	c.Status(http.StatusOK)
}
