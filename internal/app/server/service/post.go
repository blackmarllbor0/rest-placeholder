package service

import (
	"math/rand"
	"restplaceholder/internal/app/models"
	"time"

	"restplaceholder/pkg/dictionary"
)

type Posts struct {
}

func NewPosts() *Posts {
	return &Posts{}
}

func (ps Posts) GetPost() models.Post {
	p := models.Post{}
	p.ID = dictionary.NewID()
	p.Title = dictionary.NewTitle()
	p.Content = dictionary.NewContent()

	return p
}

func (ps Posts) GetPosts() models.Posts {
	rand.Seed(time.Now().UnixNano()) //nolint:staticcheck
	length := rand.Intn(50)          //nolint:gosec

	return ps.generatePosts(length)
}

func (ps Posts) GetPostsByLength(length int) models.Posts {
	return ps.generatePosts(length)
}

func (ps Posts) generatePosts(length int) models.Posts {
	posts := models.Posts{}

	for i := 0; i < length; i++ {
		posts = append(posts, ps.GetPost())
	}

	return posts
}
