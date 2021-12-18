package users

import (
	"github.com/daffashafwan/pointcuan/business/users"
	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/daffashafwan/pointcuan/controllers/user/requests"
	"github.com/daffashafwan/pointcuan/controllers/user/responses"
	"net/http"
	"github.com/labstack/echo"
)

type UserController struct {
	UserUseCase users.Usecase
}

func NewUserController(userUseCase users.Usecase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (userController UserController) Login(c echo.Context) error {

	userLogin := requests.UserLogin{}
	c.Bind(&userLogin)
	ctx := c.Request().Context()
	user, error := userController.UserUseCase.Login(ctx, userLogin.ToDomain())

	if error != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, error)
	}

	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(user))
}
