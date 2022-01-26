package admin_test
import (
	"context"
	"testing"
	"time"

	"github.com/daffashafwan/pointcuan/app/middlewares"
	"github.com/daffashafwan/pointcuan/business/admin"
	_adminMock "github.com/daffashafwan/pointcuan/business/admin/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository _adminMock.Repository
var userService admin.Usecase
var userDomain admin.Domain
var listUserDomain []admin.Domain
var token string

func setup() {
	userService = admin.NewUsecase(&userRepository, time.Second*10, middlewares.ConfigJWT{})
	userDomain = admin.Domain{
		Id:       1,
		Username: "daffashafwan111",
		Password: "$2a$12$.CA1/o7b4zddxl0JtprODOhvHb8LoOaPZ3U5HrdEUjZSznOi7d2EK",
	}
	token = "token"
	listUserDomain = append(listUserDomain, userDomain)

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
		assert.Equal(t, admin.Domain{}, data)
		assert.Error(t, err)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | Login - Error (Pass Empty)", func(t *testing.T) {
		setup()
		userDomain.Password = ""
		data, err := userService.Login(context.Background(), userDomain)
		assert.Equal(t, data.JWTToken, "")
		assert.Equal(t, admin.Domain{}, data)
		assert.Error(t, err)

		userRepository.AssertExpectations(t)
	})

	// t.Run("Test Case 3 | Login - Error (Wrong Auth)", func(t *testing.T) {
	// 	setup()
	// 	userRepository.On("UsersGetByEmail",
	// 		mock.Anything, mock.AnythingOfType("string")).Return(admin.Domain{}, businesses.ErrForTest).Once()
	// 	data, token, err := userService.Login(context.Background(), userDomain.Email, "1234")

	// 	assert.Equal(t, admin.Domain{}, data)
	// 	assert.Error(t, err)
	// 	assert.Equal(t, token, "")

	// 	userRepository.AssertExpectations(t)
	// })
}
