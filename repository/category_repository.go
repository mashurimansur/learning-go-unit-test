package repository

import "learning-go-unit-test/entity"

type CategoryRepository interface {
	FindByID(id string) *entity.Category
}
