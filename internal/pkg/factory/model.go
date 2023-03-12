package factory

import (
	"fmt"
	"restplaceholder/internal/models"
	"restplaceholder/pkg/dictionary"
)

func GetModel(model interface{}) (interface{}, error) {
	switch model {
	case models.Post{}:
		p := models.Post{}
		p.ID = dictionary.NewID()
		p.Title = dictionary.NewTitle()
		p.Content = dictionary.NewContent()

		return p, nil
	default:
		return nil, fmt.Errorf("this type is not supported")
	}
}
