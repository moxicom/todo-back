package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/moxicom/todo-back/pkg/service"
)

type Handler struct {
	service service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: *service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.CreateList)
			lists.GET("/", h.GetAllLists)
			lists.GET("/:id", h.GetList)
			lists.PUT("/:id", h.UpdateList)
			lists.DELETE("/:id", h.DeleteList)

			items := lists.Group(":id/items")
			{
				items.GET("/", h.GetAllItems)
				items.POST("/", h.CreateItem)
				items.GET("/:item_id", h.GetItem)
				items.PUT("/:items_id", h.UpdateItem)
				items.DELETE("/:items_id", h.DeleteItem)
			}
		}
	}

	return router
}
