package transaction

import (
	"net/http"
	"strconv"

	"github.com/daffashafwan/pointcuan/business/pcr_crud"
	"github.com/daffashafwan/pointcuan/business/transactions"
	"github.com/daffashafwan/pointcuan/controllers/transaction/requests"
	"github.com/daffashafwan/pointcuan/controllers/transaction/responses"
	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	TransactionUsecase transactions.Usecase
	PcrUsecase         pcrcrud.Usecase
}

func NewTransactionController(transactionUsecase transactions.Usecase, pcrUseCase pcrcrud.Usecase) *TransactionController {
	return &TransactionController{
		TransactionUsecase: transactionUsecase,
		PcrUsecase: pcrUseCase,
	}
}

func (transactionController TransactionController) Create(c echo.Context) error {
	transCreate := requests.TransactionRequest{}
	c.Bind(&transCreate)
	ctx := c.Request().Context()
	id := c.Param("id")
	convId, _ := strconv.Atoi(id)
	transCreate.UserId = convId
	pcr, err := transactionController.PcrUsecase.GetPCR(ctx)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	trs,_ := strconv.ParseFloat(transCreate.Transaction, 64)
	transCreate.Point = trs / pcr.PcrValue
	transaction, errors := transactionController.TransactionUsecase.Create(ctx, transCreate.ToDomain())
	if errors != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, errors)
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromDomain(transaction))
}

func (transactionController TransactionController) GetAll(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := transactionController.TransactionUsecase.GetAll(ctxNative)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (transactionController TransactionController) GetById(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("tid")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, errConvInt.Error())
	}
	data, err := transactionController.TransactionUsecase.GetById(ctxNative, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromDomain(data))
}

func (transactionController TransactionController) GetByUserId(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, errConvInt.Error())
	}
	data, err := transactionController.TransactionUsecase.GetByUserId(ctxNative, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (transactionController TransactionController) GetByUserIdAndStatus(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, errConvInt.Error())
	}
	transCreate := requests.TransactionRequest{}
	c.Bind(&transCreate)
	data, err := transactionController.TransactionUsecase.GetByUserIdAndStatus(ctxNative, convInt, int(transCreate.Status))
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (transactionController TransactionController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("tid")
	convId, _ := strconv.Atoi(id)
	transaction, _ := transactionController.TransactionUsecase.GetById(ctx, convId)
	transactionReq := requests.TransactionRequest{}
	var err = c.Bind(&transactionReq)
	if err != nil {
		return err
	}
	data, err := transactionController.TransactionUsecase.Update(ctx, transactionReq.ToDomain(), transaction.Id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromDomain(data))
}

func (transactionController TransactionController) Delete(c echo.Context) error {
	id := c.Param("tid")
	convInt, err := strconv.Atoi(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	err = transactionController.TransactionUsecase.Delete(ctx, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, convInt)
}
