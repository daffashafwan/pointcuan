package pcrcrud

import (
	"net/http"

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
	pcrUpdate := requests.PcrUpdate{}
	c.Bind(&pcrUpdate)
	ctx := c.Request().Context()
	data, err := pcrController.PcrUseCase.Update(ctx, pcrUpdate.ToDomain())
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (pcrController PcrController) GetPCR(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := pcrController.PcrUseCase.GetPCR(ctxNative)
	if data.Id == 0 {
		return response.ErrorResponse(c, http.StatusNotFound, err.Error())
	}
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}