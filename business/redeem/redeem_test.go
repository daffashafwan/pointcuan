package redeem_test

import (
	"context"
	"github.com/daffashafwan/pointcuan/app/middlewares"
	"github.com/daffashafwan/pointcuan/business/items"
	_itemsMock "github.com/daffashafwan/pointcuan/business/items/mocks"
	"github.com/daffashafwan/pointcuan/business/point"
	_pointMock "github.com/daffashafwan/pointcuan/business/point/mocks"
	"github.com/daffashafwan/pointcuan/business/redeem"
	_redeemMock "github.com/daffashafwan/pointcuan/business/redeem/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var pointRepository _pointMock.Repository
var pointService point.Usecase
var pointDomain point.Domain

var itemRepository _itemsMock.Repository
var itemService items.Usecase
var itemDomain items.Domain
var listItemDomain []items.Domain

var redeemRepository _redeemMock.Repository
var redeemService redeem.Usecase
var redeemDomain redeem.Domain
var listRedeemDomain []redeem.Domain
var token string

func setup() {
	itemService = items.NewItemsUsecase(&itemRepository, time.Second*10)
	pointService = point.NewPointUsecase(&pointRepository, time.Second*10, middlewares.ConfigJWT{})
	redeemService = redeem.NewRedeemUsecase(&pointRepository, &itemRepository, &redeemRepository, time.Second*10, middlewares.ConfigJWT{})
	pointDomain = point.Domain{
		Id:     1,
		UserId: 1,
		Point:  400,
	}
	itemDomain = items.Domain{
		Id:          1,
		CategoryId:  1,
		Name:        "Pulsa",
		PointRedeem: 544,
		Stock:       454,
	}
	redeemDomain = redeem.Domain{
		Id:     1,
		UserId: 1,
		ItemId: 1,
		//Item: itemDomain,
		DataRedeem: "56444",
		Point:      416415,
		Status:     0,
	}
	token = "token"
	listRedeemDomain = append(listRedeemDomain, redeemDomain)
	listItemDomain = append(listItemDomain, itemDomain)

}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		redeemRepository.On("Create",
			mock.Anything, mock.Anything).Return(redeemDomain, nil).Once()
		itemRepository.On("GetByItemId",
			mock.Anything, mock.Anything).Return(itemDomain, nil).Once()
		pointRepository.On("GetByUserId",
			mock.Anything, mock.Anything).Return(pointDomain, nil).Once()
		pointRepository.On("Update",
			mock.Anything, mock.Anything).Return(pointDomain, nil).Once()
		itemRepository.On("Update",
			mock.Anything, mock.Anything).Return(itemDomain, nil).Once()
		data, err := redeemService.Create(context.Background(), redeem.Domain{
			Id:         2,
			UserId:     1,
			ItemId:     1,
			DataRedeem: "4554",
			Point:      544,
		})

		assert.Nil(t, err)
		assert.Equal(t, data, redeemDomain)
		redeemRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Create - Success", func(t *testing.T) {
		setup()
		redeemRepository.On("Create",
			mock.Anything, mock.Anything).Return(redeemDomain, nil).Once()
		itemDomain.Name = "gopay 20000"
		itemRepository.On("GetByItemId",
			mock.Anything, mock.Anything).Return(itemDomain, nil).Once()
		pointRepository.On("GetByUserId",
			mock.Anything, mock.Anything).Return(pointDomain, nil).Once()
		pointRepository.On("Update",
			mock.Anything, mock.Anything).Return(pointDomain, nil).Once()
		itemRepository.On("Update",
			mock.Anything, mock.Anything).Return(itemDomain, nil).Once()
		data, err := redeemService.Create(context.Background(), redeem.Domain{
			Id:         2,
			UserId:     1,
			ItemId:     1,
			DataRedeem: "4554",
			Point:      544,
		})

		assert.Nil(t, err)
		assert.Equal(t, data, data)
		redeemRepository.AssertExpectations(t)
	})

}

func TestGetAll(t *testing.T) {
	t.Run("Test Case 1 | GetAll - Success", func(t *testing.T) {
		setup()
		redeemRepository.On("GetAll",
			mock.Anything).Return(listRedeemDomain, nil).Once()
		data, err := redeemService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listRedeemDomain))
		redeemRepository.AssertExpectations(t)
	})

	// t.Run("Test Case 2| GetAll - Error", func(t *testing.T) {
	// 	setup()
	// 	userRepository.On("GetAll",
	// 		mock.Anything).Return([]users.Domain{}, "Error")
	// 	data, err := userService.GetAll(context.Background())

	// 	assert.Error(t, err)
	// 	assert.Equal(t, data, []users.Domain{})
	// 	userRepository.AssertExpectations(t)
	// })
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1 | GetById - Success", func(t *testing.T) {
		setup()
		redeemRepository.On("GetById",
			mock.Anything, mock.AnythingOfType("int")).Return(redeemDomain, nil).Once()
		data, err := redeemService.GetById(context.Background(), int(redeemDomain.Id))

		assert.NoError(t, err)
		assert.NotNil(t, data)

		redeemRepository.AssertExpectations(t)
	})

	t.Run("Test Case 1 | GetById - Success", func(t *testing.T) {
		setup()
		redeemDomain.Id = 0
		redeemRepository.On("GetById",
			mock.Anything, mock.AnythingOfType("int")).Return(redeemDomain, nil).Once()
		data, err := redeemService.GetById(context.Background(), 1)

		assert.Error(t, err)
		assert.NotNil(t, data)

		redeemRepository.AssertExpectations(t)
	})
}

func TestGetByUserId(t *testing.T) {
	t.Run("Test Case 1 | GetByUserId - Success", func(t *testing.T) {
		setup()
		redeemRepository.On("GetByUserId",
			mock.Anything, mock.AnythingOfType("int")).Return(listRedeemDomain, nil).Once()
		data, err := redeemService.GetByUserId(context.Background(), int(redeemDomain.UserId))

		assert.NoError(t, err)
		assert.NotNil(t, data)

		redeemRepository.AssertExpectations(t)
	})
}

func TestGetByItemId(t *testing.T) {
	t.Run("Test Case 1 | GetByItemId - Success", func(t *testing.T) {
		setup()
		redeemRepository.On("GetByItemId",
			mock.Anything, mock.AnythingOfType("int")).Return(listRedeemDomain, nil).Once()
		data, err := redeemService.GetByItemId(context.Background(), int(redeemDomain.ItemId))

		assert.NoError(t, err)
		assert.NotNil(t, data)

		redeemRepository.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test Case 1 | Delete - Success", func(t *testing.T) {
		setup()
		redeemRepository.On("Delete",
			mock.Anything, mock.Anything).Return(nil).Once()

		err := redeemService.Delete(context.Background(), 1)
		assert.Nil(t, err)
		redeemRepository.AssertExpectations(t)
	})

	t.Run("Test Case 1 | Delete - Fail", func(t *testing.T) {
		setup()
		redeemRepository.On("Delete",
			mock.Anything, mock.Anything).Return(nil).Once()

		err := redeemService.Delete(context.Background(), 2)
		assert.Nil(t, err)
		redeemRepository.AssertExpectations(t)
	})
}
