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

}

func (h *Handler) SignIn(c *gin.Context) {

}
