package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moxicom/todo-back/models"
)

func (h *Handler) CreateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.TodoList
	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// service
	id, err := h.service.TodoList.Create(userId, input)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetAllLists(c *gin.Context) {

}

func (h *Handler) GetList(c *gin.Context) {

}

func (h *Handler) UpdateList(c *gin.Context) {

}

func (h *Handler) DeleteList(c *gin.Context) {

}
