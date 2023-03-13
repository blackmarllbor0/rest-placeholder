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

	router.GET("/post", h.GetPost)
	posts := router.Group("posts")
	{
		posts.GET("/", h.GetPosts)
		posts.GET("/:len", h.GetPostsByLength)
	}

	router.GET("/user", h.GetUser)
	users := router.Group("users")
	{
		users.GET("/", h.GetUsers)
		users.GET("/:len", h.GetUsersByLength)
	}
	return router
}
