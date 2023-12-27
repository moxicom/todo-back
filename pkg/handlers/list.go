package handlers

import (
	"net/http"
	"strconv"

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

	id, err := h.service.TodoList.Create(userId, input)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []models.TodoList `json:"data"`
}

func (h *Handler) GetAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	lists, err := h.service.TodoList.GetAll(userId)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// Router: /api/lists/:id [get]
func (h *Handler) GetList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newResponseError(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.service.TodoList.GetById(userId, listId)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateList(c *gin.Context) {

}

func (h *Handler) DeleteList(c *gin.Context) {

}
