package redeem

import (
	"context"
	"errors"
	"fmt"
	"github.com/daffashafwan/pointcuan/app/middlewares"
	"github.com/daffashafwan/pointcuan/business/items"
	"github.com/daffashafwan/pointcuan/business/point"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"strconv"
	"strings"
	"time"
)

type RedeemUsecase struct {
	PointRepo      point.Repository
	ItemRepo       items.Repository
	Repo           Repository
	contextTimeout time.Duration
	ConfigJWT      middlewares.ConfigJWT
}

func NewRedeemUsecase(pointRepo point.Repository, itemRepo items.Repository, repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &RedeemUsecase{
		ConfigJWT:      configJWT,
		PointRepo:      pointRepo,
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
	if splits[0] == "gopay" {
		mc := coreapi.Client{}
		mc.New("SB-Mid-server-u01zpyt2UKfqZ_hDTnr_Edgb", midtrans.Sandbox)
		amt, _ := strconv.ParseInt(splits[1], 0, 64)
		// 2. Initiate charge request
		chargeReq := &coreapi.ChargeReq{
			PaymentType: coreapi.PaymentTypeGopay,
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  "465164",
				GrossAmt: amt,
			},
		}

		// 3. Request to Midtrans
		coreApiRes, _ := mc.ChargeTransaction(chargeReq)
		redeem.ResponseMidtrans = coreApiRes.Actions
		fmt.Println("Response :", coreApiRes)
	}
	points,_ := tc.PointRepo.GetByUserId(ctx, domain.UserId)
	points.Point = points.Point - float64(item.PointRedeem)
	pointU, _ := tc.PointRepo.Update(ctx, points)
	fmt.Println(pointU)
	item.Stock = item.Stock - 1
	items, _ := tc.ItemRepo.Update(ctx, item)
	fmt.Println(items)
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
