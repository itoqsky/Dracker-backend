package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itoqsky/money-tracker-backend/internal/core"
)

func (h *Handler) createGroup(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var input core.Group
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	groupId, err := h.services.Group.Create(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"group_id": groupId,
	})
}

type getAllGroupsResponse struct {
	Data []core.Group `json:"data"`
}

func (h *Handler) getAllGroups(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	groups, err := h.services.Group.GetAll(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllGroupsResponse{
		Data: groups,
	})
}

func (h *Handler) getGroupById(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	group_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	group, err := h.services.Group.GetById(id, group_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, group)
}

func (h *Handler) updateGroup(c *gin.Context) {

}

func (h *Handler) deleteGroup(c *gin.Context) {

}
