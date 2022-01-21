package redeem

import (
	"context"
	"errors"
	"github.com/daffashafwan/pointcuan/app/middlewares"
	"github.com/daffashafwan/pointcuan/business/items"
	"time"
	"fmt"
	"strings"
	"strconv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type RedeemUsecase struct {
	ItemRepo       items.Repository
	Repo           Repository
	contextTimeout time.Duration
	ConfigJWT      middlewares.ConfigJWT
}

func NewRedeemUsecase(itemRepo items.Repository, repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &RedeemUsecase{
		ConfigJWT:      configJWT,
		ItemRepo:       itemRepo,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (tc *RedeemUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	if domain.DataRedeem == "" {
		return Domain{}, errors.New("data empty")
	}
	var err error

	if err != nil {
		return Domain{}, err
	}
	item, _ := tc.ItemRepo.GetByItemId(ctx, domain.ItemId)
	domain.Point = item.PointRedeem
	redeem, err := tc.Repo.Create(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}
	splits := strings.Split(item.Name, " ")
	if (splits[0] == "gopay") {
		mc := coreapi.Client{}
		mc.New("SB-Mid-server-u01zpyt2UKfqZ_hDTnr_Edgb", midtrans.Sandbox)
		amt,_ := strconv.ParseInt(splits[1],0, 64)
		// 2. Initiate charge request
		chargeReq := &coreapi.ChargeReq{
			PaymentType: coreapi.PaymentTypeGopay,
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  "1145414",
				GrossAmt: amt,
			},
		}

		// 3. Request to Midtrans
		coreApiRes, _ := mc.ChargeTransaction(chargeReq)
		redeem.ResponseMidtrans = coreApiRes.Actions
		fmt.Println("Response :", coreApiRes)
	}

	return redeem, nil
}

func (tc *RedeemUsecase) Delete(ctx context.Context, id int) error {
	err := tc.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (tc *RedeemUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	redeem, err := tc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return redeem, nil
}

func (tc *RedeemUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	redeem, err := tc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if redeem.Id == 0 {
		return Domain{}, errors.New("ID NOT FOUND")
	}
	return redeem, nil
}

func (tc *RedeemUsecase) GetByUserId(ctx context.Context, id int) ([]Domain, error) {
	redeem, err := tc.Repo.GetByUserId(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return redeem, nil
}

func (tc *RedeemUsecase) GetByItemId(ctx context.Context, id int) ([]Domain, error) {
	redeem, err := tc.Repo.GetByItemId(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return redeem, nil
}
