package categoryitem

import (
	"net/http"
	"strconv"

	categoryItems "github.com/daffashafwan/pointcuan/business/categoryItems"
	"github.com/daffashafwan/pointcuan/controllers/categoryitem/requests"
	"github.com/daffashafwan/pointcuan/controllers/categoryitem/responses"
	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/labstack/echo/v4"
)

type CategoryItemController struct {
	CategoryItemsCase categoryItems.Usecase
}

func NewCategoryController(categoryUseCase categoryItems.Usecase) *CategoryItemController {
	return &CategoryItemController{
		CategoryItemsCase: categoryUseCase,
	}
}

func (categoryItemController CategoryItemController) Create(c echo.Context) error {

	categoryItemCreate := requests.CategoryItem{}
	c.Bind(&categoryItemCreate)
	ctx := c.Request().Context()
	categoryItem, error := categoryItemController.CategoryItemsCase.Create(ctx, categoryItemCreate.ToDomain())

	if error != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, error)
	}

	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(categoryItem))
}

func (categoryItemController CategoryItemController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("cid")
	convId, _ := strconv.Atoi(id)
	catReq := requests.CategoryItem{}
	var err = c.Bind(&catReq)
	if err != nil {
		return err
	}
	data, err := categoryItemController.CategoryItemsCase.Update(ctx, catReq.ToDomain(), convId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (categoryItemController CategoryItemController) GetAll(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := categoryItemController.CategoryItemsCase.GetAll(ctxNative)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (categoryItemController CategoryItemController) GetById(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("cid")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, errConvInt.Error())
	}
	data, err := categoryItemController.CategoryItemsCase.GetById(ctxNative, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromDomain(data))
}



func (categoryItemController CategoryItemController) Delete(c echo.Context) error {
	id := c.Param("cid")
	convInt, _ := strconv.Atoi(id)
	ctx := c.Request().Context()
	var err = categoryItemController.CategoryItemsCase.Delete(ctx, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, convInt)
}