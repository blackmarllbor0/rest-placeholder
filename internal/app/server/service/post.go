package service

import (
	"math/rand"
	factory2 "restplaceholder/internal/pkg/factory"
	"time"

	"restplaceholder/internal/models"
)

type Posts struct {
}

func NewPosts() *Posts {
	return &Posts{}
}

func (ps Posts) GetPost() (models.Post, error) {
	post, err := factory2.GetModel(models.Post{})
	if err != nil {
		return models.Post{}, err
	}

	return post.(models.Post), nil
}

func (ps Posts) GetPosts() (models.Posts, error) {
	rand.Seed(time.Now().UnixNano()) //nolint:staticcheck
	length := rand.Intn(500)         //nolint:gosec

	posts, err := factory2.GetModels(models.Posts{}, length)
	if err != nil {
		return nil, err
	}

	return posts.(models.Posts), nil
}

func (ps Posts) GetPostsByLength(length int) models.Posts {
	return ps.generatePosts(length)
}

func (ps Posts) generatePosts(length int) models.Posts {
	posts := models.Posts{}

	// for i := 0; i < length; i++ {
	// 	posts = append(posts, ps.GetPost())
	// }

	return posts
}
