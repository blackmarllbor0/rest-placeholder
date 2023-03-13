package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) BlankImage(ctx *gin.Context) {
	width, err := strconv.Atoi(ctx.DefaultQuery("w", "250"))
	if err != nil {
		panic(err)
	}

	height, err := strconv.Atoi(ctx.DefaultQuery("h", "250"))
	if err != nil {
		panic(err)
	}

	clr := ctx.DefaultQuery("clr", "black")

	imgBuffer, err := h.service.GetBlankImage(width, height, clr)
	if err != nil {
		panic(err)
	}

	if _, err := ctx.Writer.Write(imgBuffer.Bytes()); err != nil {
		panic(err)
	}
}
