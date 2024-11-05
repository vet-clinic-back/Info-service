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

// // @Summary Create owner
// // @Description Create a new owner in the system
// // @Tags owners
// // @Accept json
// // @Produce json
// // @Param input body models.Owner true "owner details"
// // @Success 201 {object} models.Owner "Successfully created owner"
// // @Failure 400 {object} models.ErrorDTO "Invalid input body"
// // @Failure 409 {object} models.ErrorDTO "Owner with same email already exists"
// // @Failure 500 {object} models.ErrorDTO "Internal server error"
// // @Router /info/v1/owner/ [post]
func (h *Handler) createOwner(c *gin.Context) {
	op := "Handler.createOwner"
	log := h.log.WithField("op", op)

	var input models.Owner

	log.Debug("binding json")
	if err := c.BindJSON(&input); err != nil {
		log.Error("failed to bind json: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	log.Debug("validating input")
	if err := http_utils.ValidateCreatingOwnerDTO(input); err != nil {
		log.Error("failed to validate input: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body. all required fields must be present")
		return
	}

	log.Debug("checking if owner already exists")
	_, err := h.service.Info.GetOwner(models.Owner{
		Email: input.Email,
		Phone: input.Phone,
	})
	if err == nil {
		log.Error("failed to create new owner. Owner with same email or phone already exists")
		h.newErrorResponse(c, http.StatusConflict, "owner with same email or phone already exists")
		return
	}

	if !errors.Is(err, sql.ErrNoRows) {
		log.Error("failed to check existing owner: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to check existing owner in db")
		return
	}

	log.Debug("creating owner")
	owner, err := h.service.Info.CreateOwner(input)
	if err != nil {
		log.Error("failed to create owner: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to create owner")
		return
	}

	log.Info("successfully created owner")
	c.JSON(http.StatusCreated, owner)
}

// // @Summary Get owner
// // @Description Get owner details by ID
// // @Tags owners
// // @Produce json
// // @Param id path int true "owner ID"
// // @Success 200 {object} models.Owner "Successfully retrieved owner"
// // @Failure 404 {object} models.ErrorDTO "owner not found"
// // @Failure 500 {object} models.ErrorDTO "Internal server error"
// //  @Router /info/v1/owner/{id} [get]
func (h *Handler) getOwner(c *gin.Context) {
	op := "Handler.getOwner"
	log := h.log.WithField("op", op)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Error("invalid owner ID: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid owner ID")
		return
	}

	own := models.Owner{ID: uint(id)}

	owner, err := h.service.Info.GetOwner(own)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("owner not found: ", err.Error())
			h.newErrorResponse(c, http.StatusNotFound, "owner not found")
			return
		}
		log.Error("failed to get owner: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to get owner")
		return
	}

	log.Info("successfully retrieved owner")
	c.JSON(http.StatusOK, owner)
}

// // @Summary Get all owners
// // @Description Get all owners details
// // @Tags owners
// // @Produce json
// // @Success 200 {object} models.Owner "Successfully retrieved owners"
// // @Failure 500 {object} models.ErrorDTO "Internal server error"
// // @Router  /info/v1/owner [get]
func (h *Handler) getAllOwners(c *gin.Context) {
	op := "Handler.getAllOwners"
	log := h.log.WithField("op", op)

	log.Debug("retrieving all owners")
	owners, err := h.service.Info.GetAllOwners()
	if err != nil {
		log.Error("failed to get all owners: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to get all owners")
		return
	}

	log.Info("successfully retrieved all owners")
	c.JSON(http.StatusOK, owners)
}

// // @Summary Update owner
// // @Description Update owner details by ID
// // @Tags owners
// // @Accept json
// // @Produce json
// // @Param id path int true "owner ID"
// // @Param input body models.Owner true "owner details"
// // @Success 200 {object} models.Owner "Successfully updated owner"
// // @Failure 400 {object} models.ErrorDTO "Invalid input body or owner ID"
// // @Failure 404 {object} models.ErrorDTO "Owner not found"
// // @Failure 409 {object} models.ErrorDTO "Owner with same email already exists"
// // @Failure 500 {object} models.ErrorDTO "Internal server error"
// // @Router /info/v1/owner/{id} [put]
func (h *Handler) updateOwner(c *gin.Context) {
	op := "Handler.updateOwner"
	log := h.log.WithField("op", op)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Error("invalid owner ID: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid owner ID")
		return
	}

	var input models.Owner
	if err := c.BindJSON(&input); err != nil {
		log.Error("failed to bind json: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	input.ID = uint(id)

	// Check for existing owner with the same email (excluding the owner being updated)
	existingOwner, err := h.service.Info.GetOwner(models.Owner{Email: input.Email})
	if err == nil {
		if existingOwner.ID != input.ID { // Check if it's not the same owner
			log.Error("owner with this email already exists")
			h.newErrorResponse(c, http.StatusConflict, "owner with this email already exists")
			return
		}
	} else if !errors.Is(err, sql.ErrNoRows) {
		log.Error("failed to check existing owner: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to check existing owner in db")
		return
	}

	log.Debug("updating owner")
	updatedOwner, err := h.service.Info.UpdateOwner(input)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error("owner not found: ", err.Error())
			h.newErrorResponse(c, http.StatusNotFound, "owner not found")
			return
		}
		log.Error("failed to update owner: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to update owner")
		return
	}

	log.Info("successfully updated owner")
	c.JSON(http.StatusOK, updatedOwner)
}

// // @Summary Delete owner
// // @Description Delete owner details by ID
// // @Tags owners
// // @Accept json
// // @Produce json
// // @Param id path int true "owner ID"
// // @Param input body models.Owner true "owner details"
// // @Success 200 {object} models.Owner "Successfully deleted owner"
// // @Failure 404 {object} models.ErrorDTO "owner not found"
// // @Failure 500 {object} models.ErrorDTO "Internal server error"
// // @Router /info/v1/owner/{id} [delete]
func (h *Handler) deleteOwner(c *gin.Context) {
	op := "Handler.deleteOwner"
	log := h.log.WithField("op", op)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Error("invalid owner ID: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid owner ID")
		return
	}

	log.Debug("deleting owner")
	err = h.service.Info.DeleteOwner(uint(id))
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("owner not found: ", err.Error())
			h.newErrorResponse(c, http.StatusNotFound, "owner not found")
			return
		}
		log.Error("failed to delete owner: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to delete owner")
		return
	}

	log.Info("successfully deleted owner")
	c.Status(http.StatusOK)
}
