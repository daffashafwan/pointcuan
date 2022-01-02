package categoryitem

import (
	"net/http"

	categoryItems "github.com/daffashafwan/pointcuan/business/categoryItem"
	"github.com/daffashafwan/pointcuan/controllers/categoryItem/responses"
	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/labstack/echo/v4"
)

type CategoryItemController struct {
	CategoryItemsCase categoryItems.Usecase
}




func (categoryItemController CategoryItemController) Create(c echo.Context) error {

	categoryItemCreate := requests.categoryItemCreate{}
	c.Bind(&categoryItemCreate)
	ctx := c.Request().Context()
	categoryItem, error := categoryItemController.CategoryItemUseCase.Create(ctx, categoryItemCreate.ToDomain())

	if error != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, error)
	}

	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(categoryItem))
}

func (categoryItemController CategoryItemController) Update(c echo.Context) error {

}

func (categoryItemController CategoryItemController) GetAll(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := categoryitemController.CategoryItemUseCase.GetAll(ctxNative)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (categoryItemController CategoryItemController) GetById(c echo.Context) error {
	
}



func (categoryItemController CategoryItemController) Delete(c echo.Context) error {
	
}