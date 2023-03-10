package service

import (
	"restplaceholder/internal/models"
	"restplaceholder/pkg/dictionary"
)

func (s Service) GetPost() models.Post {
	p := models.Post{}
	p.ID = dictionary.NewId()
	p.Title = dictionary.NewTitle()
	p.Content = dictionary.NewContent()

	return p
}

func (s Service) GetPosts(length int) models.Posts {
	posts := models.Posts{}

	if length > 500 {
		panic("не думаю, чтобы вам могло понадобиться так много постов")
	}

	if length < 2 {
		return models.Posts{s.GetPost()}
	}

	for i := 0; i < length; i++ {
		posts = append(posts, s.GetPost())
	}

	return posts
}
