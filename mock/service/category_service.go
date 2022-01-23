package service

import (
	"errors"

	"github.com/pranotobudi/go-with-tests/mock/entity"
	"github.com/pranotobudi/go-with-tests/mock/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category not found")
	}
	return category, nil
}
