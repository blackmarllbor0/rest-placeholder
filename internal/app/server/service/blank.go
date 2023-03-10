package service

import (
	"bytes"

	"restplaceholder/pkg/img"
)

func (s Service) GetBlankImage(width, height int, clr string) (*bytes.Buffer, error) {
	return img.GenerateBlankImage(width, height, clr)
}
