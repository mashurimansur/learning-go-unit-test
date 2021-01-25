package service

import (
	"errors"
	"learning-go-unit-test/entity"
	"learning-go-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service *CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindByID(id)

	if category == nil {
		return nil, errors.New("Category not found")
	} else {
		return category, nil
	}
}
