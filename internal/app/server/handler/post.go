package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetPost(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.GetPost())
}

func (h Handler) GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.GetPosts())
}

func (h Handler) GetPostsByLength(c *gin.Context) {
	length, err := strconv.Atoi(c.Param("len"))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, h.service.GetPostsByLength(length))
}