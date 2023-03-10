package handler

import (
	"github.com/gin-gonic/gin"

	"restplaceholder/internal/app/server/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/blank", h.BlankImage)

	posts := router.Group("post")
	{
		posts.GET("/", h.GetPost)
		posts.GET("/:len", h.GetPosts)
	}

	return router
}
