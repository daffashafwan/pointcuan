package users

import (
	"net/http"
	"strconv"

	"github.com/daffashafwan/pointcuan/business/users"
	"github.com/daffashafwan/pointcuan/controllers/user/requests"
	"github.com/daffashafwan/pointcuan/controllers/user/responses"
	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/labstack/echo/v4"
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

func (userController UserController) GetById(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	data, err := userController.UserUseCase.GetById(ctxNative, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (userController UserController) Verif(c echo.Context) error {
	ctxNative := c.Request().Context()
	token := c.Param("token")
	data, err := userController.UserUseCase.GetByToken(ctxNative, token)
	if data.Status == "1"{
		return response.SuccessResponse(c,http.StatusOK, "Anda Sudah Pernah Melakukan Verifikasi")
	}
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	userVerif := requests.UserRegister{
		Name: data.Name,
		Username: data.Username,
		Status: "1",
		Password: data.Password,
		Email: data.Email,
		Address: data.Address,
		Token: data.Token,
	}
	err = c.Bind(&userVerif)
	ctx := c.Request().Context()
	data, errs := userController.UserUseCase.Verif(ctx, userVerif.ToDomain(), data.Id)
	if errs != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (userController UserController) Delete(c echo.Context) error {
	id := c.Param("id")
	convInt, err := strconv.Atoi(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = userController.UserUseCase.Delete(ctx, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c,http.StatusOK, convInt)
}

func (userController UserController) ForgotPassword(c echo.Context) error {

	userForgotPassword := requests.UserForgotPassword{}
	c.Bind(&userForgotPassword)
	ctx := c.Request().Context()
	user, error := userController.UserUseCase.Create(ctx, userForgotPassword.ToDomain())

	if error != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, error)
	}

	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(user))
}