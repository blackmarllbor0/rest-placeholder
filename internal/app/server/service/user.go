package service

import (
	"math/rand"
	"restplaceholder/internal/app/models"
	"restplaceholder/internal/app/server/service/interfaces"
	"restplaceholder/pkg/dictionary"
	"time"
)

type User struct {
	interfaces.Posts
}

func NewUser(posts interfaces.Posts) *User {
	return &User{Posts: posts}
}

func (u User) GetUser() models.User {
	user := u.generateUser()
	user.Friends = u.GetUsers()

	return user
}

func (u User) GetUsers() models.Users {
	rand.Seed(time.Now().UnixNano()) //nolint:staticcheck
	length := rand.Intn(10)          //nolint:gosec

	return u.generateUsers(length)
}

func (u User) GetUsersByLength(length int) models.Users {
	return u.generateUsers(length)
}

func (u User) generateUsers(length int) models.Users {
	users := models.Users{}

	for i := 0; i < length; i++ {
		users = append(users, u.generateUser())
	}

	return users
}

func (u User) generateUser() models.User {
	user := models.User{}
	user.ID = dictionary.NewID()
	user.Name = dictionary.NewTitle()
	user.Nick = dictionary.NewTitle()
	user.Posts = u.GetPosts()

	return user
}
