package redeem

import (
	"net/http"
	"strconv"
	"github.com/daffashafwan/pointcuan/business/items"
	"github.com/daffashafwan/pointcuan/business/redeem"
	"github.com/daffashafwan/pointcuan/controllers/redeem/requests"
	"github.com/daffashafwan/pointcuan/controllers/redeem/responses"

	//"github.com/daffashafwan/pointcuan/helpers/randomizer"
	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/labstack/echo/v4"
)

type RedeemController struct {
	RedeemUsecase redeem.Usecase
	ItemsUsecase  items.Usecase
}

func NewRedeemController(redeemUsecase redeem.Usecase, itemsUsecase items.Usecase) *RedeemController {
	return &RedeemController{
		RedeemUsecase: redeemUsecase,
		ItemsUsecase:  itemsUsecase,
	}
}

func (redeemController RedeemController) Create(c echo.Context) error {
	transCreate := requests.RedeemRequest{}
	c.Bind(&transCreate)
	ctx := c.Request().Context()
	id := c.Param("id")
	convId, _ := strconv.Atoi(id)
	transCreate.UserId = convId
	// transCreate.RefId = randomizer.Randomize(3)
	//item, _ := redeemController.ItemsUsecase.GetByItemId(ctx, transCreate.ItemId)
	transaction, errors := redeemController.RedeemUsecase.Create(ctx, transCreate.ToDomain())
	if errors != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, errors)
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromDomain(transaction))
}

func (redeemController RedeemController) GetAll(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := redeemController.RedeemUsecase.GetAll(ctxNative)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (redeemController RedeemController) GetById(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("rid")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, errConvInt.Error())
	}
	data, err := redeemController.RedeemUsecase.GetById(ctxNative, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromDomain(data))
}

func (redeemController RedeemController) GetByUserId(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, errConvInt.Error())
	}
	data, err := redeemController.RedeemUsecase.GetByUserId(ctxNative, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (redeemController RedeemController) GetByItemId(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("iid")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, errConvInt.Error())
	}
	data, err := redeemController.RedeemUsecase.GetByItemId(ctxNative, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (redeemController RedeemController) Delete(c echo.Context) error {
	id := c.Param("id")
	convInt, err := strconv.Atoi(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	err = redeemController.RedeemUsecase.Delete(ctx, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, convInt)
}
