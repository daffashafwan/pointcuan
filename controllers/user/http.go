package users

import (
	"github.com/daffashafwan/pointcuan/business/users"
	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/daffashafwan/pointcuan/controllers/user/requests"
	"github.com/daffashafwan/pointcuan/controllers/user/responses"
	"net/http"
	"github.com/labstack/echo"
	"strconv"
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

func (userController UserController) Register(c echo.Context) error {

	userRegister := requests.UserRegister{}
	c.Bind(&userRegister)
	ctx := c.Request().Context()
	user, error := userController.UserUseCase.Create(ctx, userRegister.ToDomain())

	if error != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, error)
	}

	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(user))
}

func (userController UserController) Update(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)
	}
	userRegister := requests.UserRegister{}
	err = c.Bind(&userRegister)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := userController.UserUseCase.Update(ctx, userRegister.ToDomain(), convId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (userController UserController) GetAll(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := userController.UserUseCase.GetAll(ctxNative)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

