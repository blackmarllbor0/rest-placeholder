package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h Handler) GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, h.service.GetUser())
}

func (h Handler) GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, h.service.GetUsers())
}

func (h Handler) GetUsersByLength(c *gin.Context) {
	length, err := strconv.Atoi(c.Param("len"))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, h.service.GetUsersByLength(length))
}
