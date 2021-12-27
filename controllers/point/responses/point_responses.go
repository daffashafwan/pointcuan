package responses

import (
	"github.com/daffashafwan/pointcuan/business/point"
)

type PointResponse struct {
	UserId int     `json:"userId"`
	Point  float64 `json:"point"`
}

func FromDomain(domain point.Domain) PointResponse {
	return PointResponse{
		UserId: domain.UserId,
		Point:  domain.Point,
	}
}
