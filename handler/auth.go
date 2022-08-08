package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	storage "github.com/mskydream/audio-cloud/storage"
)

func (h *Handler) signUp(c *gin.Context) {
	var user storage.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.CreateUser(user)

	if errors.Is(err, storage.UserExists) {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type refreshTokensInput struct {
	RefreshTooken string `json:"refresh_token" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	userId, token, err := h.services.Authorization.GenerateAccessToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := h.services.Authorization.GenerateRefreshToken(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokensResponse{
		Token:        token,
		RefreshToken: refreshToken,
	})
}

func (h *Handler) refreshTokens(c *gin.Context) {
	var input refreshTokensInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	userId, newRefreshToken, err := h.services.UpdateRefreshToken(input.RefreshTooken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newAccessToken, err := h.services.UpdateAccessToken(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":         newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
