package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itoqsky/money-tracker-backend/internal/core"
)

//	@Summary		Create group
//	@Security		ApiKeyAuth
//	@Tags			group
//	@Description	create group
//	@ID				create-group
//	@Accept			json
//	@Produce		json
//	@Param			input	body		core.Group	true	"group info"
//	@Success		200		{object}	string		"group_id"
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/groups [post]

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

//	@Summary		Get all groups
//	@Security		ApiKeyAuth
//	@Tags			group
//	@Description	get all groups
//	@ID				get-all-groups
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	getAllGroupsResponse
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/groups [get]

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

	c.JSON(http.StatusOK, getAllGroupsResponse{groups})
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

//	@Summary		Update group
//	@Security		ApiKeyAuth
//	@Tags			group
//	@Description	update group
//	@ID				update-group
//	@Accept			json
//	@Produce		json
//	@Param			input	body		core.UpdateGroupInput	true	"group info"
//	@Success		200		{object}	string					"status"
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/groups/{id} [put]

func (h *Handler) updateGroup(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	group_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input core.UpdateGroupInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.services.Group.Update(id, group_id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

//	@Summary		Delete group
//	@Security		ApiKeyAuth
//	@Tags			group
//	@Description	delete group
//	@ID				delete-group
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	string	"status"
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/groups/{id} [delete]

func (h *Handler) deleteGroup(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	group_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Group.Delete(id, group_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
