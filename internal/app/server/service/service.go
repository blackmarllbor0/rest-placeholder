package service

import (
	"restplaceholder/internal/app/server/service/interfaces"
)

type Service struct {
	interfaces.BlankImage
	interfaces.Posts
}

func NewService() *Service {
	return &Service{}
}
