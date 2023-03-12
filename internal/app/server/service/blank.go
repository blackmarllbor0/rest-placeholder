package service

import (
	"bytes"

	"restplaceholder/pkg/img"
)

type Blank struct {
}

func NewBlank() *Blank {
	return &Blank{}
}

func (b Blank) GetBlankImage(width, height int, clr string) (*bytes.Buffer, error) {
	return img.GenerateBlankImage(width, height, clr)
}
