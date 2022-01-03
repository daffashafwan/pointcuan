package admins

import (
	"net/http"
	"strconv"

	admins "github.com/daffashafwan/pointcuan/business/admin"
	"github.com/daffashafwan/pointcuan/controllers/admin/requests"
	"github.com/daffashafwan/pointcuan/controllers/admin/responses"
	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/labstack/echo/v4"
)

type AdminController struct {
	AdminUseCase admins.Usecase
}

func NewAdminController(adminUseCase admins.Usecase) *AdminController {
	return &AdminController{
		AdminUseCase: adminUseCase,
	}
}

func (adminController AdminController) Login(c echo.Context) error {

	adminLogin := requests.AdminLogin{}
	c.Bind(&adminLogin)
	ctx := c.Request().Context()
	admin, errors := adminController.AdminUseCase.Login(ctx, adminLogin.ToDomain())

	if errors != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, errors)
	}

	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(admin))
}

func (adminController AdminController) GetById(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	data, err := adminController.AdminUseCase.GetById(ctxNative, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}