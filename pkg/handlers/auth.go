package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moxicom/todo-back/models"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	// h.Service.Auth.CreateUser...
	id, err := h.service.Auth.CreateUser(input)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {

}
