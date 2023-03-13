package service

import (
	"restplaceholder/internal/app/server/service/interfaces"
)

type Service struct {
	interfaces.BlankImage
	interfaces.Posts
	interfaces.Users
}

func NewService() *Service {
	s := Service{}
	s.BlankImage = NewBlank()
	s.Posts = NewPosts()
	s.Users = NewUser(s.Posts)

	return &s
}
