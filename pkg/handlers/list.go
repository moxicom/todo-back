package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateList(c *gin.Context) {
	id, _ := c.Get(userCtx) // skip ok because we check its existence in middleware

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
