package pcrcrud_test

import (
	"context"
	pcrcrud "github.com/daffashafwan/pointcuan/business/pcr_crud"
	_categoryMock "github.com/daffashafwan/pointcuan/business/pcr_crud/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var userRepository _categoryMock.Repository
var userService pcrcrud.Usecase
var userDomain pcrcrud.Domain
var listUserDomain []pcrcrud.Domain
var token string

func setup() {
	userService = pcrcrud.NewPcrcase(&userRepository, time.Second*10)
	userDomain = pcrcrud.Domain{
		Id:       1,
		PcrValue: 800,
	}
	token = "token"
	listUserDomain = append(listUserDomain, userDomain)

}

func TestGetAll(t *testing.T) {
	t.Run("Test Case 1 | GetPcr - Success", func(t *testing.T) {
		setup()
		userRepository.On("GetPCR",
			mock.Anything).Return(userDomain, nil).Once()
		data, err := userService.GetPCR(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, data, userDomain)
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

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1 | Update - Success", func(t *testing.T) {
		setup()
		userRepository.On("GetPCR",
			mock.Anything).Return(userDomain, nil).Once()
		userRepository.On("Update",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		data, err := userService.Update(context.Background(), pcrcrud.Domain{
			Id:       1,
			PcrValue: 9000,
		})

		assert.Nil(t, err)
		assert.Equal(t, data, userDomain)
		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 1 | Update - Success", func(t *testing.T) {
		setup()
		userDomain.Id = 0
		userRepository.On("GetPCR",
			mock.Anything).Return(userDomain, nil).Once()
		userRepository.On("Create",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		data, err := userService.Update(context.Background(), pcrcrud.Domain{
			Id:       1,
			PcrValue: 9000,
		})

		assert.Nil(t, err)
		assert.Equal(t, data, userDomain)
		userRepository.AssertExpectations(t)
	})

}
