package items_test

import (
	"context"
	"github.com/daffashafwan/pointcuan/business/items"
	_itemsMock "github.com/daffashafwan/pointcuan/business/items/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var itemRepository _itemsMock.Repository
var itemService items.Usecase
var itemDomain items.Domain
var listItemDomain []items.Domain

func setup() {
	itemService = items.NewItemsUsecase(&itemRepository, time.Second*10)
	itemDomain = items.Domain{
		Id:          1,
		CategoryId:  1,
		Name:        "Pulsa",
		PointRedeem: 544,
		Stock:       454,
	}
	listItemDomain = append(listItemDomain, itemDomain)
}

func TestGetAll(t *testing.T) {
	t.Run("Test Case 1 | GetAll - Success", func(t *testing.T) {
		setup()
		itemRepository.On("GetAll",
			mock.Anything).Return(listItemDomain, nil).Once()
		data, err := itemService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listItemDomain))
		itemRepository.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1 | GetById - Success", func(t *testing.T) {
		setup()
		itemRepository.On("GetByItemId",
			mock.Anything, mock.AnythingOfType("int")).Return(itemDomain, nil).Once()
		data, err := itemService.GetByItemId(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, data)

		itemRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | GetById - Error - User Id 0", func(t *testing.T) {
		setup()
		itemDomain.Id = 0
		itemRepository.On("GetByItemId",
			mock.Anything, mock.AnythingOfType("int")).Return(itemDomain, nil).Once()
		data, err := itemService.GetByItemId(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, data, data)

		itemRepository.AssertExpectations(t)
	})
}

func TestGetByCategoryId(t *testing.T) {
	t.Run("Test Case 1 | GetByCategoryId - Success", func(t *testing.T) {
		setup()
		itemRepository.On("GetByCategoryId",
			mock.Anything, mock.Anything).Return(listItemDomain, nil).Once()
		data, err := itemService.GetByCategoryId(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listItemDomain))
		itemRepository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		itemRepository.On("Create",
			mock.Anything, mock.Anything).Return(itemDomain, nil).Once()

		data, err := itemService.Create(context.Background(), items.Domain{
			Id:          1,
			CategoryId:  1,
			Name:        "Pulsa",
			PointRedeem: 544,
			Stock:       454,
		})

		assert.Nil(t, err)
		assert.Equal(t, data, itemDomain)
		itemRepository.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1 | Update - Success", func(t *testing.T) {
		setup()
		itemRepository.On("Update",
			mock.Anything, mock.Anything).Return(itemDomain, nil).Once()

		data, err := itemService.Update(context.Background(), items.Domain{
			Id:       1,
			CategoryId:  1,
			Name:        "Pulsa",
			PointRedeem: 544,
			Stock:       454,
		}, 1)

		assert.Nil(t, err)
		assert.Equal(t, data, itemDomain)
		itemRepository.AssertExpectations(t)
	})

}

func TestUpdateStock(t *testing.T) {
	t.Run("Test Case 1 | UpdateStock - Success", func(t *testing.T) {
		setup()
		itemRepository.On("UpdateStock",
			mock.Anything, mock.Anything).Return(itemDomain, nil).Once()
		itemRepository.On("GetByItemId",
			mock.Anything, mock.Anything).Return(itemDomain, nil).Once()
		data, err := itemService.UpdateStock(context.Background(), items.Domain{
			Id:       1,
			CategoryId:  1,
			Name:        "Pulsa",
			PointRedeem: 544,
			Stock:       454,
		}, 1)

		assert.Nil(t, err)
		assert.Equal(t, data, itemDomain)
		itemRepository.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	t.Run("Test Case 1 | Delete - Success", func(t *testing.T) {
		setup()
		itemRepository.On("Delete",
			mock.Anything, mock.Anything).Return(nil).Once()

		err := itemService.Delete(context.Background(), 1)
		assert.Nil(t, err)
		itemRepository.AssertExpectations(t)
	})

	t.Run("Test Case 1 | Delete - Fail", func(t *testing.T) {
		setup()
		itemRepository.On("Delete",
			mock.Anything, mock.Anything).Return(nil).Once()

		err := itemService.Delete(context.Background(), 2)
		assert.Nil(t, err)
		itemRepository.AssertExpectations(t)
	})
}
