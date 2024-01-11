package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moxicom/todo-back/models"
)

// @Summary Create item
// @Security ApiKeyAuth
// @Tags item
// @Description create item
// @ID create-item
// @Accept  json
// @Produce  json
// @Param listId path int true "list id"
// @Param input body models.Item true "Item"
// @Success 200 {object} responseId
// @Failure 400,404 {object} errorMsg
// @Failure 500 {object} errorMsg
// @Failure default {object} errorMsg
// @Router /api/lists/{listId}/items [post]
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
	c.JSON(http.StatusOK, responseId{
		Id: itemId,
	})
}

type getAllItemsResponse struct {
	Data []models.Item `json:"data"`
}

// @Summary Get all items
// @Security ApiKeyAuth
// @Tags item
// @Description Get all items
// @ID get-all-items
// @Accept  json
// @Produce  json
// @Param listId path int true "list id"
// @Success 200 {object} getAllItemsResponse
// @Failure 400,404 {object} errorMsg
// @Failure 500 {object} errorMsg
// @Failure default {object} errorMsg
// @Router /api/lists/{listId}/items/ [get]
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

// @Summary Get item by id
// @Security ApiKeyAuth
// @Tags item
// @Description Get item by id
// @ID get-item-by-id
// @Accept  json
// @Produce  json
// @Param listId path int true "list id"
// @Param itemId path int true "item id"
// @Success 200 {object} models.Item
// @Failure 400,404 {object} errorMsg
// @Failure 500 {object} errorMsg
// @Failure default {object} errorMsg
// @Router /api/lists/{listId}/items/{itemId} [get]
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

// @Summary Update item by id
// @Security ApiKeyAuth
// @Tags item
// @Description Update item by id
// @ID update-item-by-id
// @Accept  json
// @Produce  json
// @Param listId path int true "list id"
// @Param itemId path int true "item id"
// @Param input body models.Item true "Item content"
// @Success 200 {object} responseId
// @Failure 400,404 {object} errorMsg
// @Failure 500 {object} errorMsg
// @Failure default {object} errorMsg
// @Router /api/lists/{listId}/items/{itemId} [put]
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

	c.JSON(http.StatusOK, responseId{
		Id: itemId,
	})
}

// @Summary Delete item by id
// @Security ApiKeyAuth
// @Tags item
// @Description Delete item by id
// @ID delete-item-by-id
// @Accept  json
// @Produce  json
// @Param listId path int true "list id"
// @Param itemId path int true "item id"
// @Success 200 {object} responseId
// @Failure 400,404 {object} errorMsg
// @Failure 500 {object} errorMsg
// @Failure default {object} errorMsg
// @Router /api/lists/{listId}/items/{itemId} [delete]
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

	c.JSON(http.StatusOK, responseId{
		Id: itemId,
	})
}
