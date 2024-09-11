package service

import (
	"golang-unit-test/entity"
	"golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}

	// mock on get 3
	categoryRepository.Mock.On("FindById", "3").Return(category)

	// get on mock FindById 3
	result, err := categoryService.Get("3")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
