package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/itoqsky/money-tracker-backend/internal/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/itoqsky/money-tracker-backend/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	routes := gin.New()

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := routes.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := routes.Group("/api", h.userIdentity)
	{
		groups := api.Group("/groups")
		{
			groups.POST("/", h.createGroup)
			groups.GET("/", h.getAllGroups)
			groups.GET("/:id", h.getGroupById)
			groups.PUT("/:id", h.updateGroup)
			groups.DELETE("/:id", h.deleteGroup)

			users := groups.Group(":id/users")
			{
				users.POST("/", h.inviteUser)
				users.GET("/", h.getAllUsers)
				users.DELETE("/", h.kickUser)
			}

			purchases := groups.Group(":id/purchases")
			{
				purchases.POST("/", h.createPurchase)
				purchases.GET("/", h.getAllPurchases)
				purchases.GET("/:p_id", h.getPurchaseById)
				purchases.PUT("/:p_id", h.updatePurchase)
				purchases.DELETE("/:p_id", h.deletePurchase)
			}
		}

		debts := api.Group("/debts")
		{
			debts.GET("/", h.getAllDebts)
			debts.PUT("/", h.updateDebt)
		}
	}

	return routes
}
