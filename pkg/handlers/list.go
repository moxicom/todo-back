package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moxicom/todo-back/models"
)

// @Summary Create todo list
// @Security ApiKeyAuth
// @Tags lists
// @Description create todo list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body models.TodoList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorMsg
// @Failure 500 {object} errorMsg
// @Failure default {object} errorMsg
// @Router /api/lists [post]
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

// @Summary Get All Lists
// @Security ApiKeyAuth
// @Tags lists
// @Description get all lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListsResponse
// @Failure 400,404 {object} errorMsg
// @Failure 500 {object} errorMsg
// @Failure default {object} errorMsg
// @Router /api/lists [get]
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

// @Summary Get List By Id
// @Security ApiKeyAuth
// @Tags lists
// @Description get list by id
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "list id"
// @Success 200 {object} models.TodoList
// @Failure 400,404 {object} errorMsg
// @Failure 500 {object} errorMsg
// @Failure default {object} errorMsg
// @Router /api/lists/{id} [get]
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

// @Summary Update list
// @Security ApiKeyAuth
// @Tags lists
// @Description update list by id
// @ID update-list
// @Accept  json
// @Produce  json
// @Param id path int true "list id"
// @Param input body models.TodoList true "update list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorMsg
// @Failure 500 {object} errorMsg
// @Failure default {object} errorMsg
// @Router /api/lists/{id} [put]
func (h *Handler) UpdateList(c *gin.Context) {
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

	var input models.TodoList
	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.TodoList.Update(userId, listId, input)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": listId,
	})
}

// @Summary Delete list
// @Security ApiKeyAuth
// @Tags lists
// @Description delete list by id
// @ID delete-list
// @Accept  json
// @Produce  json
// @Success 200
// @Param id path int true "list id"
// @Failure 400,404 {object} errorMsg
// @Failure 500 {object} errorMsg
// @Failure default {object} errorMsg
// @Router /api/lists/{id} [delete]
func (h *Handler) DeleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.TodoList.Delete(userId, listId)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}
