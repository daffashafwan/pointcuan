package responses

import (
	"time"

	categoryItems "github.com/daffashafwan/pointcuan/business/categoryItems"
)

type CategoryItemResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Svg       string    `json:"svg"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain categoryItems.Domain) CategoryItemResponse {
	return CategoryItemResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Svg:       domain.Svg,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(domain []categoryItems.Domain) []CategoryItemResponse {
	var list []CategoryItemResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
