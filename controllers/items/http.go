package users

import (
	"net/http"
	"strconv"
	"github.com/daffashafwan/pointcuan/business/items"
	"github.com/daffashafwan/pointcuan/controllers/items/requests"
	"github.com/daffashafwan/pointcuan/controllers/items/responses"
	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/labstack/echo/v4"
)

type ItemsController struct {
	ItemsUsecase items.Usecase
}

func NewItemsController(itemsUsecase items.Usecase) *ItemsController {
	return &ItemsController{
		ItemsUsecase: itemsUsecase,
	}
}

func (itemsController ItemsController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	items, _ := itemsController.ItemsUsecase.GetByItemId(ctx, convId)
	itemsRequest := requests.ItemRequest{}
	err = c.Bind(&itemsRequest)
	if err != nil {
		return err
	}
	data, err := itemsController.ItemsUsecase.Update(ctx, itemsRequest.ToDomain(), items.Id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (itemsController ItemsController) Delete(c echo.Context) error {
	id := c.Param("id")
	convInt, err := strconv.Atoi(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	err = itemsController.ItemsUsecase.Delete(ctx, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, convInt)
}