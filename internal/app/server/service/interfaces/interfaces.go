package interfaces

import (
	"bytes"
	"restplaceholder/internal/app/models"
)

type BlankImage interface {
	GetBlankImage(width, height int, clr string) (buffer *bytes.Buffer, err error)
}

type Posts interface {
	GetPost() (models.Post, error)
	GetPosts() (models.Posts, error)
	GetPostsByLength(length int) models.Posts
}

type Users interface {
	GetUser() models.User
	GetUsers() models.Users
	GetUsersByLength(length int) models.Users
}
