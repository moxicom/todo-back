package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	// headers
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		newResponseError(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	headerFields := strings.Split(header, " ")
	if len(headerFields) != 2 || headerFields[0] != "Bearer" {
		newResponseError(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	// Bearer <token>
	if len(headerFields[1]) == 0 {
		newResponseError(c, http.StatusUnauthorized, "token is empty")
	}

	userId, err := h.service.Auth.ParseToken(headerFields[1])
	if err != nil {
		newResponseError(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
