package pcrcrud

import (
	"net/http"
	"strconv"

	pcrcrud "github.com/daffashafwan/pointcuan/business/pcr_crud"
	"github.com/daffashafwan/pointcuan/controllers/pcr_crud/requests"
	"github.com/daffashafwan/pointcuan/controllers/pcr_crud/responses"
	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/labstack/echo/v4"
)

type PcrController struct {
	PcrUseCase pcrcrud.Usecase
}

func NewPcrController(pcrUseCase pcrcrud.Usecase) *PcrController {
	return &PcrController{
		PcrUseCase: pcrUseCase,
	}
}
func (pcrController PcrController) Update(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)
	}
	pcrUpdate := requests.PcrUpdate{}
	err = c.Bind(&pcrUpdate)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := pcrController.PcrUseCase.Update(ctx, pcrUpdate.ToDomain(), convId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (pcrController PcrController) GetById(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	data, err := pcrController.PcrUseCase.GetById(ctxNative, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}