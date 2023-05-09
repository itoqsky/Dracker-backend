package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itoqsky/money-tracker-backend/internal/core"
)

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

type InputgetPurchaseById struct {
	PurchaseId int `json:"purchase_id" binding:"required"`
}

func (h *Handler) getPurchaseById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	var input InputgetPurchaseById
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	purchace, err := h.services.Purchase.GetById(input.PurchaseId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, purchace)
}

func (h *Handler) updatePurchase(c *gin.Context) {
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
	input.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Purchase.Update(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

type InputDeletePurchase struct {
	GroupId int `json:"group_id" binding:"required"`
}

func (h *Handler) deletePurchase(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	purchaseId, err := strconv.Atoi(c.Param("purchase_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input InputDeletePurchase
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	purhcase := core.Purchase{
		ID:      purchaseId,
		BuyerId: id,
		GroupId: input.GroupId,
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
