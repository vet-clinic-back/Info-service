package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vet-clinic-back/info-service/internal/models"
)

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, models.ErrorDTO{message})
}
