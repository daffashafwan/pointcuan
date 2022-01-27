package transactions_test

import (
	"context"
	"github.com/daffashafwan/pointcuan/app/middlewares"
	"testing"
	"time"
	"github.com/daffashafwan/pointcuan/business/users"
	_usersMock "github.com/daffashafwan/pointcuan/business/users/mocks"

	"github.com/daffashafwan/pointcuan/business/point"
	_pointMock "github.com/daffashafwan/pointcuan/business/point/mocks"

	"github.com/daffashafwan/pointcuan/business/transactions"
	_transactionMock "github.com/daffashafwan/pointcuan/business/transactions/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository _usersMock.Repository
var userService users.Usecase
var userDomain users.Domain
var listUserDomain []users.Domain
var token string

var pointRepository _pointMock.Repository
var pointService point.Usecase
var pointDomain point.Domain

var transactionRepository _transactionMock.Repository
var transactionService transactions.Usecase
var transactionDomain transactions.Domain
var listTransactionDomain []transactions.Domain

func setup() {
	userService = users.NewUserUsecase(&userRepository, time.Second*10, middlewares.ConfigJWT{})
	pointService = point.NewPointUsecase(&pointRepository, time.Second*10, middlewares.ConfigJWT{})
	transactionService = transactions.NewTransactionUsecase(&transactionRepository, &pointRepository, time.Second*10, middlewares.ConfigJWT{})
	pointDomain = point.Domain{
		Id:     1,
		UserId: 1,
		Point:  500,
	}
	userDomain = users.Domain{
		Id:       1,
		Username: "daffashafwan111",
		Password: "$2a$12$.CA1/o7b4zddxl0JtprODOhvHb8LoOaPZ3U5HrdEUjZSznOi7d2EK",
		Email:    "daffashafwan@mantap.io",
		Address:  "Malang",
		Token:    "123456",
	}
	transactionDomain = transactions.Domain{
		Id:                    1,
		UserId:                1,
		TransactionDate:       time.Now(),
		Transaction:           622222,
		TransactionAttachment: "attachment",
		Status:                2,
		Point:                 5000,
		Description:           "gas",
	}
	token = "token"
	listUserDomain = append(listUserDomain, userDomain)
	listTransactionDomain = append(listTransactionDomain, transactionDomain)
}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		transactionRepository.On("Create",
			mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		data, err := transactionService.Create(context.Background(), transactionDomain)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, data, transactionDomain)
		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Create - Success", func(t *testing.T) {
		setup()
		transactionRepository.On("Failed",
			mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		transactionDomain.TransactionAttachment = ""
		data, err := transactionService.Create(context.Background(), transactionDomain)

		assert.Error(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, data, data)
		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | Create - Failed", func(t *testing.T) {
		setup()
		transactionRepository.On("Create",
			mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		transactionDomain.Transaction = 0
		data, err := transactionService.Create(context.Background(), transactionDomain)

		assert.Error(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, data, data)
		userRepository.AssertExpectations(t)
	})
}
func TestGetAll(t *testing.T) {
	t.Run("Test Case 1 | GetAll - Success", func(t *testing.T) {
		setup()
		transactionRepository.On("GetAll",
			mock.Anything).Return(listTransactionDomain, nil).Once()
		data, err := transactionService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listTransactionDomain))
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

func TestGeyByUserIdAndStatus(t *testing.T) {
	t.Run("Test Case 1 | GetAll - Success", func(t *testing.T) {
		setup()
		transactionRepository.On("GetByUserIdAndStatus",
			mock.Anything, mock.Anything, mock.Anything).Return(listTransactionDomain, nil).Once()
		data, err := transactionService.GetByUserIdAndStatus(context.Background(), 1, 2)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listTransactionDomain))
		userRepository.AssertExpectations(t)
	})
}

func TestGetByUserId(t *testing.T) {
	t.Run("Test Case 1 | GetAll - Success", func(t *testing.T) {
		setup()
		transactionRepository.On("GetByUserId",
			mock.Anything, mock.Anything, mock.Anything).Return(listTransactionDomain, nil).Once()
		data, err := transactionService.GetByUserId(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listTransactionDomain))
		userRepository.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1 | GetById - Success", func(t *testing.T) {
		setup()
		transactionRepository.On("GetById",
			mock.Anything, mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		data, err := transactionService.GetById(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, data, transactionDomain)
		userRepository.AssertExpectations(t)
	})
}

// func TestDelete(t *testing.T) {
// 	t.Run("Test Case 1 | Delete - Success", func(t *testing.T) {
// 		setup()
// 		transactionRepository.On("Delete",
// 			mock.Anything, mock.Anything).Return(nil).Once()

// 		err := transactionService.Delete(context.Background(), 1)
// 		assert.Nil(t, err)
// 		transactionRepository.AssertExpectations(t)
// 	})
// }

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1 | Update - Success", func(t *testing.T) {
		setup()
		pointRepository.On("GetByUserId",
			mock.Anything, mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		transactionRepository.On("Update",
			mock.Anything, mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		data, err := transactionService.Update(context.Background(), transactions.Domain{
			Id:                    1,
			UserId:                1,
			TransactionDate:       time.Now(),
			Transaction:           622222,
			TransactionAttachment: "attachment",
			Status:                1,
			Point:                 5000,
			Description:           "gas",
		}, 1)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, data, transactionDomain)
		userRepository.AssertExpectations(t)
	})

	// t.Run("Test Case 2 | Update - Success", func(t *testing.T) {
	// 	setup()
	// 	transactionRepository.On("Update",
	// 		mock.Anything, mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
	// 	pointRepository.On("GetByUserId",
	// 		mock.Anything, mock.Anything).Return(pointDomain, nil).Once()
	// 	pointRepository.On("Update",
	// 		mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(pointDomain, nil).Once()
		
	// 	data, err := transactionService.Update(context.Background(), transactions.Domain{
	// 		Id:                    1,
	// 		UserId:                1,
	// 		TransactionDate:       time.Now(),
	// 		Transaction:           622222,
	// 		TransactionAttachment: "attachment",
	// 		Status:                2,
	// 		Point:                 5000,
	// 		Description:           "gas",
	// 	}, 1)
	// 	fmt.Println(err)
	// 	assert.NoError(t, err)
	// 	assert.NotNil(t, data)
	// 	assert.Equal(t, data, data)
	// 	userRepository.AssertExpectations(t)
	// })
}
