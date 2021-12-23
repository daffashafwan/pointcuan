package users

import (
	"net/http"
	"strconv"
	"github.com/daffashafwan/pointcuan/business/point"
	"github.com/daffashafwan/pointcuan/controllers/point/requests"
	"github.com/daffashafwan/pointcuan/controllers/point/responses"
	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/labstack/echo/v4"
)

type PointController struct {
	PointUsecase point.Usecase
}

func NewPointController(pointUsecase point.Usecase) *PointController {
	return &PointController{
		PointUsecase: pointUsecase,
	}
}

func (pointController PointController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)
	}
	point, _ := pointController.PointUsecase.GetByUserId(ctx, convId)
	pointReq := requests.PointRequest{}
	err = c.Bind(&pointReq)
	if err != nil {
		return err
	}
	data, err := pointController.PointUsecase.Update(ctx, pointReq.ToDomain(), point.Id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (pointController PointController) Delete(c echo.Context) error {
	id := c.Param("id")
	convInt, err := strconv.Atoi(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = pointController.PointUsecase.Delete(ctx, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c,http.StatusOK, convInt)
}