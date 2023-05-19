package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itoqsky/money-tracker-backend/internal/core"
)

//	@Summary		Get all debts
//	@Security		ApiKeyAuth
//	@Tags			debt
//	@Description	get all debts
//	@ID				get-all-debts
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	getAllDebtsResponse
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/debts [get]

type getAllDebtsResponse struct {
	Debts   []core.Debt `json:"debts"`
	Credits []core.Debt `json:"credits"`
}

func (h *Handler) getAllDebts(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	debts, credits, err := h.services.Debt.GetAll(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllDebtsResponse{debts, credits})
}

//	@Summary		Update debt
//	@Security		ApiKeyAuth
//	@Tags			debt
//	@Description	update debt
//	@ID				update-debt
//	@Accept			json
//	@Produce		json
//	@Param			input	body		core.Debt	true	"debt info"
//	@Success		200		{object}	string		"status"
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/debts [put]

func (h *Handler) updateDebt(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var debt core.Debt
	if err := c.BindJSON(&debt); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	debt.CreditorID = id

	err = h.services.Debt.Update(debt)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// func (h *Handler) getDebtById(c *gin.Context) {
// }

// func (h *Handler) createDebt(c *gin.Context) {

// }

// func (h *Handler) deleteDebt(c *gin.Context) {

// }
