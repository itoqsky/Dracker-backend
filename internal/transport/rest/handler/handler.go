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
			groups.POST("/", h.CreateGroup)
			groups.GET("/", h.GetAllGroups)
			groups.GET("/:id", h.GetGroupById)
			groups.PUT("/:id", h.UpdateGroup)
			groups.DELETE("/:id", h.DeleteGroup)

			users := groups.Group("/:id/users")
			{
				users.GET("/", h.GetAllUsers)
				users.POST("/", h.CreateUser)
				users.DELETE("/:user_id", h.DeleteUser)
			}

			purchases := api.Group(":id/purchases")
			{
				purchases.GET("/", h.GetAllPurchases)
				purchases.GET("/:purchase_id", h.GetPurchaseById)
				purchases.POST("/", h.CreatePurchase)
				purchases.PUT("/:purchase_id", h.UpdatePurchase)
				purchases.DELETE("/:purchase_id", h.DeletePurchase)
			}

			debts := api.Group(":id/debts")
			{
				debts.GET("/", h.GetAllDebts)
				debts.GET("/:debt_id", h.GetDebtById)
				debts.PUT("/:debt_id", h.UpdateDebt)
			}
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
