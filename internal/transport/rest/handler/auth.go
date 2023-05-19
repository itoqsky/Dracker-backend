package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itoqsky/money-tracker-backend/internal/core"
)

//	@Summary		Sign up
//	@Tags			auth
//	@Description	sign up
//	@ID				sign-up
//	@Accept			json
//	@Produce		json
//	@Param			input	body		core.User	true	"user info"
//	@Success		200		{integer}	integer		1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/auth/sign-up [post]

func (h *Handler) signUp(c *gin.Context) {
	var input core.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//	@Summary		Sign in
//	@Tags			auth
//	@Description	sign in
//	@ID				sign-in
//	@Accept			json
//	@Produce		json
//	@Param			input	body		signInInput	true	"credentials"
//	@Success		200		{object}	string		"token"
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/auth/sign-in [post]

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
