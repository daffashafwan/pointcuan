package faq

import (
	"net/http"
	"strconv"

	FAQ "github.com/daffashafwan/pointcuan/business/FAQ"
	"github.com/daffashafwan/pointcuan/controllers/faq/requests"
	"github.com/daffashafwan/pointcuan/controllers/faq/responses"
	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/labstack/echo/v4"
)

type FaqController struct {
	FaqCase FAQ.Usecase
}

func NewFAQController(faqUseCase FAQ.Usecase) *FaqController {
	return &FaqController{
		FaqCase: faqUseCase,
	}
}

func (faqController FaqController) Create(c echo.Context) error {

	faqCreate := requests.FaqRequest{}
	c.Bind(&faqCreate)
	ctx := c.Request().Context()
	faq, error := faqController.FaqCase.Create(ctx, faqCreate.ToDomain())

	if error != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, error)
	}

	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(faq))
}

func (faqController FaqController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("fid")
	convId, _ := strconv.Atoi(id)
	catReq := requests.FaqRequest{}
	var err = c.Bind(&catReq)
	if err != nil {
		return err
	}
	data, err := faqController.FaqCase.Update(ctx, catReq.ToDomain(), convId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (faqController FaqController) GetAll(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := faqController.FaqCase.GetAll(ctxNative)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (faqController FaqController) GetActive(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := faqController.FaqCase.GetActive(ctxNative)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}



func (faqController FaqController) Delete(c echo.Context) error {
	id := c.Param("fid")
	convInt, _ := strconv.Atoi(id)
	ctx := c.Request().Context()
	var err = faqController.FaqCase.Delete(ctx, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, convInt)
}