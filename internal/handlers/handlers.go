package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vet-clinic-back/info-service/docs"
	"github.com/vet-clinic-back/info-service/internal/logging"
	"github.com/vet-clinic-back/info-service/internal/service"
)

type Handler struct {
	log     *logging.Logger
	service *service.Service
}

func NewHandler(log *logging.Logger, service *service.Service) *Handler {
	return &Handler{log: log, service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	info := router.Group("/info")
	{
		v1 := info.Group("/v1")
		{
			pets := v1.Group("/pets")
			{
				pets.POST("/", h.createPet)
				pets.GET("/", h.getPets)
				pets.GET("/:id", h.getPet)
				pets.PUT("/:id", h.updatePet)
				pets.DELETE("/:id", h.deletePet)
			}
			medCard := v1.Group("/record")
			{
				entries := medCard.Group("/entries")
				{
					entries.GET("/", h.getEntries)
					entries.POST("/", h.createEntry)
					entries.DELETE("/", func(context *gin.Context) {})
				}
			}
			//owner := v1.Group("/owner")
			//{
			//	owner.POST("/", h.createOwner)
			//	owner.GET("/:id", h.getOwner)
			//	owner.GET("/", h.getAllOwners)
			//	owner.PUT("/:id", h.updateOwner)
			//	owner.DELETE("/:id", h.deleteOwner)
			//}
		}
	}

	return router
}
