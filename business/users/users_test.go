package users_test

import (
	"context"
	"testing"
	"time"

	"github.com/daffashafwan/pointcuan/app/middlewares"
	"github.com/daffashafwan/pointcuan/business/users"
	_usersMock "github.com/daffashafwan/pointcuan/business/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository _usersMock.Repository
var userService users.Usecase
var userDomain users.Domain
var listUserDomain []users.Domain
var token string

func setup() {
	userService = users.NewUserUsecase(&userRepository, time.Second*10, middlewares.ConfigJWT{})
	userDomain = users.Domain{
		Id:       1,
		Username: "daffashafwan111",
		Password: "$2a$12$.CA1/o7b4zddxl0JtprODOhvHb8LoOaPZ3U5HrdEUjZSznOi7d2EK",
		Email:    "daffashafwan@mantap.io",
		Address:  "Malang",
		Token:    "123456",
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
		assert.Equal(t, data, users.Domain{})

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | GetById - Error", func(t *testing.T) {
		setup()
		userRepository.On("GetById",
			mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, nil).Once()
		data, err := userService.GetById(context.Background(), int(userDomain.Id))

		assert.Error(t, err)
		assert.Equal(t, data, users.Domain{})

		userRepository.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Test Case 1 | Login - Success", func(t *testing.T) {
		setup()
		userRepository.On("Login",
			mock.Anything, mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		data, err := userService.Login(context.Background(), userDomain)

		data.JWTToken = ""
		assert.NoError(t, err)
		assert.Equal(t, data, userDomain)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Login - Error (Username Empty)", func(t *testing.T) {
		setup()
		userDomain.Username = ""
		data, err := userService.Login(context.Background(), userDomain)
		assert.Equal(t, data.JWTToken, "")
		assert.Equal(t, users.Domain{}, data)
		assert.Error(t, err)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | Login - Error (Pass Empty)", func(t *testing.T) {
		setup()
		userDomain.Password = ""
		data, err := userService.Login(context.Background(), userDomain)
		assert.Equal(t, data.JWTToken, "")
		assert.Equal(t, users.Domain{}, data)
		assert.Error(t, err)

		userRepository.AssertExpectations(t)
	})

	// t.Run("Test Case 3 | Login - Error (Wrong Auth)", func(t *testing.T) {
	// 	setup()
	// 	userRepository.On("UsersGetByEmail",
	// 		mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, businesses.ErrForTest).Once()
	// 	data, token, err := userService.Login(context.Background(), userDomain.Email, "1234")

	// 	assert.Equal(t, users.Domain{}, data)
	// 	assert.Error(t, err)
	// 	assert.Equal(t, token, "")

	// 	userRepository.AssertExpectations(t)
	// })
}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		userRepository.On("GetByUsername",
			mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, nil).Once()
		userRepository.On("GetByEmail",
			mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, nil).Once()
		userRepository.On("Create",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		data, err := userService.Create(context.Background(), users.Domain{
			Name:     "Nama",
			Password: "123",
			Username: "daffashafwan",
			Email:    "user@tesss.io",
			Address:  "Indonesia",
		})

		assert.Nil(t, err)
		assert.Equal(t, data, userDomain)
		userRepository.AssertExpectations(t)
	})

}

func TestGetByToken(t *testing.T) {
	t.Run("Test Case 1 | GetByToken - Success", func(t *testing.T) {
		setup()
		userRepository.On("GetByToken",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		data, err := userService.GetByToken(context.Background(), "123456")
		assert.Nil(t, err)
		assert.Equal(t, data, data)
		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | GetByToken - ID Not Found", func(t *testing.T) {
		setup()
		userRepository.On("GetByToken",
			mock.Anything, mock.Anything).Return(users.Domain{}, nil).Once()

		data, err := userService.GetByToken(context.Background(), "dada")
		assert.Error(t, err)
		assert.Equal(t, data, data)
		userRepository.AssertExpectations(t)
	})
}

func TestVerify(t *testing.T) {
	t.Run("Test Case 1 | Verify - Success", func(t *testing.T) {
		setup()
		userRepository.On("Update",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		data, err := userService.Verify(context.Background(), users.Domain{
			Id:       1,
			Username: "daffashafwan111",
			Password: "$2a$12$.CA1/o7b4zddxl0JtprODOhvHb8LoOaPZ3U5HrdEUjZSznOi7d2EK",
			Email:    "daffashafwan@mantap.io",
			Address:  "Malang",
			Token:    "123456",
		}, 1)
		assert.Nil(t, err)
		assert.Equal(t, data, data)
		userRepository.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1 | Update - Success", func(t *testing.T) {
		setup()
		userRepository.On("Update",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		data, err := userService.Update(context.Background(), users.Domain{
			Id:       1,
			Username: "daffashafwan121",
			Password: "$2a$12$.CA1/o7b4zddxl0JtprODOhvHb8LoOaPZ3U5HrdEUjZSznOi7d2EK",
			Email:    "daffashafwan@mantap.io",
			Address:  "Malang",
			Token:    "123456",
		}, 1)
		assert.Nil(t, err)
		assert.Equal(t, data, data)
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

func TestForgotPassword(t *testing.T) {
	t.Run("Test Case 1 | Forgot Password - Success", func(t *testing.T) {
		setup()
		userRepository.On("GetByEmail",
			mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, nil).Once()
		userRepository.On("Update",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		data, err := userService.ForgotPassword(context.Background(), "daffashafwan@mantap.io")

		assert.Nil(t, err)
		assert.Equal(t, data, userDomain)
		userRepository.AssertExpectations(t)
	})

}

func TestResetPassword(t *testing.T) {
	t.Run("Test Case 1 | ResetPassword - Success", func(t *testing.T) {
		setup()
		userRepository.On("GetById",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		userRepository.On("Update",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		data, err := userService.ResetPassword(context.Background(), "1234567", "1234567", 1)

		assert.Nil(t, err)
		assert.Equal(t, data, userDomain)
		userRepository.AssertExpectations(t)
	})

	// t.Run("Test Case 1 | ResetPassword - Fail", func(t *testing.T) {
	// 	setup()
	// 	userRepository.On("GetById",
	// 		mock.Anything, mock.Anything).Return(userDomain, nil).Once()
	// 	userRepository.On("Update",
	// 		mock.Anything, mock.Anything).Return(userDomain, nil).Once()

	// 	data, err := userService.ResetPassword(context.Background(), "1234567", "1511514", 1)
	// 	fmt.Println(data)
	// 	assert.Error(t, err)
	// 	assert.Equal(t, data, users.Domain{})
	// 	userRepository.AssertExpectations(t)
	// })

}
