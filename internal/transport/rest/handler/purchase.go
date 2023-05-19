package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itoqsky/money-tracker-backend/internal/core"
)

//	@Summary		Create purchase
//	@Security		ApiKeyAuth
//	@Tags			purchase
//	@Description	create purchase
//	@ID				create-purchase
//	@Accept			json
//	@Produce		json
//	@Param			input	body		core.Purchase	true	"purchase info"
//	@Success		200		{object}	string			"purchase_id"
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/groups/{id}/purchases [post]

func (h *Handler) createPurchase(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var input core.Purchase
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.BuyerId = id
	input.GroupId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.services.Purchase.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

type getAllPurchasesResponse struct {
	Data []core.Purchase `json:"data"`
}

//	@Summary		Get all purchases
//	@Security		ApiKeyAuth
//	@Tags			purchase
//	@Description	get all purchases
//	@ID				get-all-purchases
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	getAllPurchasesResponse
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/groups/{id}/purchases [get]

func (h *Handler) getAllPurchases(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	groupId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	purchases, err := h.services.Purchase.GetAll(groupId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllPurchasesResponse{purchases})
}

func (h *Handler) getPurchaseById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("p_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	purchace, err := h.services.Purchase.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, purchace)
}

//	@Summary		Update purchase
//	@Security		ApiKeyAuth
//	@Tags			purchase
//	@Description	update purchase
//	@ID				update-purchase
//	@Accept			json
//	@Produce		json
//	@Param			input	body		core.Purchase	true	"purchase info"
//	@Success		200		{object}	string			"status"
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/groups/{id}/purchases/{p_id} [put]

func (h *Handler) updatePurchase(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var purchace core.Purchase
	if err := c.BindJSON(&purchace); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	purchace.BuyerId = id
	purchace.GroupId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	purchace.ID, err = strconv.Atoi(c.Param("p_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Purchase.Update(purchace)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

//	@Summary		Delete purchase
//	@Security		ApiKeyAuth
//	@Tags			purchase
//	@Description	delete purchase
//	@ID				delete-purchase
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	string	"status"
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/groups/{id}/purchases/{p_id} [delete]

func (h *Handler) deletePurchase(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	groupId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	purchaseId, err := strconv.Atoi(c.Param("p_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	purhcase := core.Purchase{
		ID:      purchaseId,
		BuyerId: id,
		GroupId: groupId,
	}

	err = h.services.Purchase.Delete(purhcase)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
