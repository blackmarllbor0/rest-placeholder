package interfaces

import (
	"bytes"

	"restplaceholder/internal/models"
)

type BlankImage interface {
	GetBlankImage(width, height int, clr string) (buffer *bytes.Buffer, err error)
}

type Posts interface {
	GetPost() models.Post
	GetPosts() models.Posts
	GetPostsByLength(length int) models.Posts
}
