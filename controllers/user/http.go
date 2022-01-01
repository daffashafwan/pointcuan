package users

import (
	"net/http"
	"strconv"

	"github.com/daffashafwan/pointcuan/business/point"
	"github.com/daffashafwan/pointcuan/business/users"
	pointRequest "github.com/daffashafwan/pointcuan/controllers/point/requests"
	"github.com/daffashafwan/pointcuan/controllers/user/requests"
	"github.com/daffashafwan/pointcuan/controllers/user/responses"
	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.Usecase
	PointUsecase point.Usecase
}

func NewUserController(userUseCase users.Usecase, pointUsecase point.Usecase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
		PointUsecase: pointUsecase,
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
	point, errs := userController.PointUsecase.GetByUserId(ctxNative, data.Id)

	data.Point = point.Point
	if errs != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (userController UserController) Verify(c echo.Context) error {
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
		Id: data.Id,
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
	data, errs := userController.UserUseCase.Verify(ctx, userVerif.ToDomain(), data.Id)
	if errs != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	pointReq := pointRequest.PointRequest{
		UserId: data.Id,
		Point: 0,
	}
	point, errors := userController.PointUsecase.Create(ctx, pointReq.ToDomain())
	data.Point = point.Point
	if errors != nil {
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
	var err error
	userForgot := requests.UserForgotPassword{}
	c.Bind(&userForgot)
	ctx := c.Request().Context()
	users, err := userController.UserUseCase.ForgotPassword(ctx, userForgot.Email)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(users))
}

func (userController UserController) VerifyTokenPassword(c echo.Context) error {
	ctxNative := c.Request().Context()
	token := c.Param("token")
	data, err := userController.UserUseCase.GetByToken(ctxNative, token)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (userController UserController) ResetPassword(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)
	}
	userReset := requests.UserResetPassword{}
	err = c.Bind(&userReset)
	if err != nil || userReset.Password != userReset.RetypePassword {
		return response.ErrorResponse(c, http.StatusNotAcceptable, err)
	}
	ctx := c.Request().Context()
	data, err := userController.UserUseCase.ResetPassword(ctx, userReset.Password, convId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}
