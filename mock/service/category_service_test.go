package service

import (
	"testing"

	"github.com/pranotobudi/go-with-tests/mock/entity"
	"github.com/pranotobudi/go-with-tests/mock/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
// var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// Arrange
	var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
	var categoryService = CategoryService{Repository: categoryRepository}
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	// Act
	// program mock
	category, err := categoryService.Get("1")

	// Assert
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_Found(t *testing.T) {
	// Arrange
	var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
	var categoryService = CategoryService{Repository: categoryRepository}

	category := entity.Category{
		ID:   "1",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)

	// Act
	result, err := categoryService.Get("2")

	// Assert
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, category.ID, result.ID)
	assert.Equal(t, category.Name, result.Name)
}
