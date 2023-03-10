package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) BlankImage(c *gin.Context) {
	width, err := strconv.Atoi(c.DefaultQuery("w", "250"))
	if err != nil {
		panic(err)
	}

	height, err := strconv.Atoi(c.DefaultQuery("h", "250"))
	if err != nil {
		panic(err)
	}

	clr := c.DefaultQuery("clr", "black")

	imgBuffer, err := h.service.GetBlankImage(width, height, clr)
	if err != nil {
		panic(err)
	}

	if _, err := c.Writer.Write(imgBuffer.Bytes()); err != nil {
		panic(err)
	}
}
