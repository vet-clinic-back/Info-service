package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vet-clinic-back/info-service/internal/models"
	"github.com/vet-clinic-back/info-service/internal/utils"
)

// @Summary Create Pet
// @Description Create a new pet in the system
// @Tags pets
// @Accept json
// @Produce json
// @Param input body models.Pet true "Pet details"
// @Success 201 {object} models.Pet "Successfully created pet"
// @Failure 400 {object} models.ErrorDTO "Invalid input body"
// @Failure 500 {object} models.ErrorDTO "Internal server error"
// TODO// @Router /pet/create [post]
func (h *Handler) createPet(c *gin.Context) {
	op := "Handler.createPet"
	log := h.log.WithField("op", op)

	var input models.Pet

	log.Debug("binding json")
	if err := c.BindJSON(&input); err != nil {
		log.Error("failed to bind json: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	log.Debug("validating input")
	if err := utils.ValidateCreatingPetDTO(input); err != nil {
		log.Error("failed to validate input: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body. all required fields must be present")
		return
	}

	log.Debug("creating pet")
	pet, err := h.service.Info.CreatePet(input)
	if err != nil {
		log.Error("failed to create pet: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to create pet")
		return
	}

	log.Info("successfully created pet")
	c.JSON(http.StatusCreated, pet)
}

// @Summary Get Pet
// @Description Get pet details by ID
// @Tags pets
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {object} models.Pet "Successfully retrieved pet"
// @Failure 404 {object} models.ErrorDTO "Pet not found"
// @Failure 500 {object} models.ErrorDTO "Internal server error"
// @Router /pet/get/{id} [get]
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
		if err == sql.ErrNoRows {
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
// @Tags pets
// @Produce json
// @Param
// @Success 200 {object} models.Pet "Successfully retrieved pets"
// @Failure 500 {object} models.ErrorDTO "Internal server error"
// TODO// @Router  /pet/get [get]
func (h *Handler) getAllPets(c *gin.Context) {
	op := "Handler.getAllPets"
	log := h.log.WithField("op", op)

	log.Debug("retrieving all pets")
	pets, err := h.service.Info.GetAllPets()
	if err != nil {
		log.Error("failed to get all pets: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to get all pets")
		return
	}

	log.Info("successfully retrieved all pets")
	c.JSON(http.StatusOK, pets)
}

// @Summary Update Pet
// @Description Update pet details by ID
// @Tags pets
// @Accept json
// @Produce json
// @Param id path int true "Pet ID"
// @Param input body models.Pet true "Pet details"
// @Success 200 {object} models.Pet "Successfully updated pet"
// @Failure 400 {object} models.ErrorDTO "Invalid input body or pet ID"
// @Failure 404 {object} models.ErrorDTO "Pet not found"
// @Failure 500 {object} models.ErrorDTO "Internal server error"
// @Router /pet/update/{id} [put]
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
		if err == sql.ErrNoRows {
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
// @Tags pets
// @Accept json
// @Produce json
// @Param id path int true "Pet ID"
// @Param input body models.Pet true "Pet details"
// @Success 200 {object} models.Pet "Successfully deleted pet"
// @Failure 404 {object} models.ErrorDTO "Pet not found"
// @Failure 500 {object} models.ErrorDTO "Internal server error"
// @Router /pet/delete/{id} [delete]
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
	err = h.service.Info.DeletePet(uint(id))
	if err != nil {
		if err == sql.ErrNoRows {
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
