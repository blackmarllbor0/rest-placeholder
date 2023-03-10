package handler

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetPost(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.GetPost())
}

func (h Handler) GetPosts(c *gin.Context) {
	length, err := strconv.Atoi(c.Param("len"))
	if err != nil {
		panic(err)
	}

	if length == 0 {
		rand.Seed(time.Now().UnixNano())
		length = rand.Intn(500)
	}

	c.JSON(http.StatusOK, h.service.GetPosts(length))
}
