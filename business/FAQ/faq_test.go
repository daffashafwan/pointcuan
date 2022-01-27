package FAQ_test

import (
	"context"
	"github.com/daffashafwan/pointcuan/app/middlewares"
	"github.com/daffashafwan/pointcuan/business/FAQ"
	_categoryMock "github.com/daffashafwan/pointcuan/business/FAQ/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var userRepository _categoryMock.Repository
var userService FAQ.Usecase
var userDomain FAQ.Domain
var listUserDomain []FAQ.Domain
var token string

func setup() {
	userService = FAQ.NewFAQUsecase(&userRepository, time.Second*10, middlewares.ConfigJWT{})
	userDomain = FAQ.Domain{
		Id:       1,
		Question: "Siapa tim terkuat",
		Answer:   "Real Madrid",
		Status:   0,
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

func TestGetActive(t *testing.T) {
	t.Run("Test Case 1 | GetActive - Success", func(t *testing.T) {
		setup()
		userRepository.On("GetActive",
			mock.Anything).Return(listUserDomain, nil).Once()
		data, err := userService.GetActive(context.Background())

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

func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		userRepository.On("Create",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		data, err := userService.Create(context.Background(), FAQ.Domain{
			Question: "Siapa tim terkuat",
			Answer:   "Real Madrid",
			Status:   0,
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

		data, err := userService.Update(context.Background(), FAQ.Domain{
			Id:       1,
			Question: "Siapa tim terkuat",
			Answer:   "Real Madrid",
			Status:   0,
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

	t.Run("Test Case 1 | Delete - Fail", func(t *testing.T) {
		setup()
		userRepository.On("Delete",
			mock.Anything, mock.Anything).Return(nil).Once()

		err := userService.Delete(context.Background(), 2)
		assert.Nil(t, err)
		userRepository.AssertExpectations(t)
	})
}
