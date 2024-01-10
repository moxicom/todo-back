package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moxicom/todo-back/models"
)

// @Summary CreateItem
// @Tags auth
// @Description sign in to account
// @ID sign-in
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorMsg
// @Failure 500 {object} errorMsg
// @Failure default {object} errorMsg
// @Router /auth/sign-in [post]
func (h *Handler) CreateItem(c *gin.Context) {
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

	var input models.Item
	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := h.service.Item.Create(userId, listId, input)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": itemId,
	})
}

type getAllItemsResponse struct {
	Data []models.Item `json:"data"`
}

func (h *Handler) GetAllItems(c *gin.Context) {
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

	items, err := h.service.Item.GetAll(userId, listId)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}

func (h *Handler) GetItem(c *gin.Context) {
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

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	item, err := h.service.Item.GetById(userId, listId, itemId)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) UpdateItem(c *gin.Context) {
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

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.Item.Update(userId, listId, itemId, input)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": itemId})
}

func (h *Handler) DeleteItem(c *gin.Context) {
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

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.Item.Delete(userId, listId, itemId)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": itemId})
}
