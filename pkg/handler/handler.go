package handler

import (
	"github.com/alexandraya/mydairy-backend/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		tasks := api.Group("/tasks")
		{
			tasks.POST("/", h.createTask)
			tasks.GET("/:date", h.getTasksByDate)
			tasks.GET("/:id", h.getTaskById)
			tasks.PUT("/:id", h.updateTask)
			tasks.DELETE("/:id", h.deleteTask)
		}
	}

	return router
}
