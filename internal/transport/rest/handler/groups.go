package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateGroup(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetGroupById(c *gin.Context) {

}

func (h *Handler) GetAllGroups(c *gin.Context) {

}

func (h *Handler) UpdateGroup(c *gin.Context) {

}

func (h *Handler) DeleteGroup(c *gin.Context) {

}
