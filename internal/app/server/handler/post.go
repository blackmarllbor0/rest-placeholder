package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetPost(ctx *gin.Context) {
	post, err := h.service.GetPost()
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, post)
}

func (h Handler) GetPosts(c *gin.Context) {
	posts, err := h.service.GetPosts()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, posts)
}

func (h Handler) GetPostsByLength(c *gin.Context) {
	length, err := strconv.Atoi(c.Param("len"))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, h.service.GetPostsByLength(length))
}
