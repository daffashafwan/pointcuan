package categoryItems_test

import (
	"context"
	"github.com/daffashafwan/pointcuan/app/middlewares"
	"github.com/daffashafwan/pointcuan/business/categoryItems"
	_categoryMock "github.com/daffashafwan/pointcuan/business/categoryItems/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var userRepository _categoryMock.Repository
var userService categoryItems.Usecase
var userDomain categoryItems.Domain
var listUserDomain []categoryItems.Domain
var token string

func setup() {
	userService = categoryItems.NewCategoryUsecase(&userRepository, time.Second*10, middlewares.ConfigJWT{})
	userDomain = categoryItems.Domain{
		Id:       1,
		Name: "PLN",
	}
	token = "token"
	listUserDomain = append(listUserDomain, userDomain)

}

func TestGetAll(t *testing.T) {
	t.Run("Test Case 1 | GetAll - Success", func(t *testing.T) {
		setup()
		userRepository.On("GetAll",
			mock.Anything).Return(listUserDomain, nil).Once()
		data, err := userService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listUserDomain))
		userRepository.AssertExpectations(t)
	})

	// t.Run("Test Case 2| GetAll - Error", func(t *testing.T) {
	// 	setup()
	// 	userRepository.On("GetAll",
	// 		mock.Anything).Return([]users.Domain{}, "Error")
	// 	data, err := userService.GetAll(context.Background())

	// 	assert.Error(t, err)
	// 	assert.Equal(t, data, []users.Domain{})
	// 	userRepository.AssertExpectations(t)
	// })
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1 | GetById - Success", func(t *testing.T) {
		setup()
		userRepository.On("GetById",
			mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		data, err := userService.GetById(context.Background(), int(userDomain.Id))

		assert.NoError(t, err)
		assert.NotNil(t, data)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | GetById - Error - User Id 0", func(t *testing.T) {
		setup()
		userDomain.Id = 0
		userRepository.On("GetById",
			mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		data, err := userService.GetById(context.Background(), 7)

		assert.Error(t, err)
		assert.Equal(t, data, categoryItems.Domain{})

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | GetById - Error", func(t *testing.T) {
		setup()
		userRepository.On("GetById",
			mock.Anything, mock.AnythingOfType("int")).Return(categoryItems.Domain{}, nil).Once()
		data, err := userService.GetById(context.Background(), int(userDomain.Id))

		assert.Error(t, err)
		assert.Equal(t, data, categoryItems.Domain{})

		userRepository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		userRepository.On("Create",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		data, err := userService.Create(context.Background(), categoryItems.Domain{
			Name:     "Nama",
		})

		assert.Nil(t, err)
		assert.Equal(t, data, userDomain)
		userRepository.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1 | Update - Success", func(t *testing.T) {
		setup()
		userRepository.On("Update",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		data, err := userService.Update(context.Background(), categoryItems.Domain{
			Id: 1,
			Name:     "Nama",
		}, 1)

		assert.Nil(t, err)
		assert.Equal(t, data, userDomain)
		userRepository.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	t.Run("Test Case 1 | Delete - Success", func(t *testing.T) {
		setup()
		userRepository.On("Delete",
			mock.Anything, mock.Anything).Return(nil).Once()

		err := userService.Delete(context.Background(), 1)
		assert.Nil(t, err)
		userRepository.AssertExpectations(t)
	})
}