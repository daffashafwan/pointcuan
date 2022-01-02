package responese

import (
	"time"

	categoryItems "github.com/daffashafwan/pointcuan/business/categoryItem"
)

type CategoryItemResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain category.Domain) CategoryItemResponse {
	return CategoryItemsResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(domain []categoryItems.Domain) []CategoryItemsResponse {
	var list []CategoryItemsResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
