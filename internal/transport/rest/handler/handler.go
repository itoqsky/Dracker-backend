package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/itoqsky/money-tracker-backend/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	routes := gin.New()

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
				users.GET("/", h.getAllUsers)
				users.POST("/", h.inviteUser)
				users.DELETE("/", h.kickUser)
			}

			purchases := api.Group(":id/purchases")
			{
				purchases.GET("/", h.getAllPurchases)
				purchases.GET("/:purchase_id", h.getPurchaseById)
				purchases.POST("/", h.createPurchase)
				purchases.PUT("/:purchase_id", h.updatePurchase)
				purchases.DELETE("/:purchase_id", h.deletePurchase)
			}

		}

		debts := api.Group("/debts")
		{
			debts.GET("/", h.getAllDebts)
			debts.GET("/:id", h.getDebtById)
			debts.PUT("/:id", h.updateDebt)
		}

		// users := api.Group("/users")
		// {
		// 	users.GET("/")
		// 	users.GET("/:id")
		// 	users.POST("/")
		// 	users.PUT("/:id")
		// 	users.DELETE("/:id")
		// }
	}

	return routes
}
