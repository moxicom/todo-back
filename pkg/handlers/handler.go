package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/moxicom/todo-back/docs"
	"github.com/moxicom/todo-back/pkg/service"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:     true,
		AllowMethods:        []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:        []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:       []string{"Content-Length"},
		AllowCredentials:    true,
		MaxAge:              12 * 3600,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}

	api := router.Group("/api", h.userIdentity)
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
				items.PUT("/:item_id", h.UpdateItem)
				items.DELETE("/:item_id", h.DeleteItem)
			}
		}
	}

	return router
}
