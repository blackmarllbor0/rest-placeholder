package factory

import (
	"fmt"

	"restplaceholder/internal/app/server/service"
	"restplaceholder/internal/models"
)

func GetModels(model interface{}, length int) (interface{}, error) {
	switch model {
	case models.Posts{}: //nolint:typecheck
		postM := models.Posts{}
		postT, err := service.Posts{}.GetPost()
		if err != nil {
			return nil, err
		}

		res := generateModel(postM, postT, length)

		return res, nil
	default:
		return nil, fmt.Errorf("this type is not supported")
	}
}

func generateModel(models interface{}, model interface{}, length int) interface{} {
	m := models.([]interface{})

	for i := 0; i < length; i++ {
		m = append(m, model)
	}

	return m
}
