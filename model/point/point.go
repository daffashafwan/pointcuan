package point

import (
	"github.com/daffashafwan/pointcuan/business/point"
	"time"
)

type Point struct {
	Id        int `gorm:"primaryKey]"`
	UserId    int `gorm:"unique"`
	Point     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (po *Point) ToDomain() point.Domain {
	return point.Domain{
		Id:        po.Id,
		UserId:    po.UserId,
		Point:     po.Point,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
	}
}

func ToListDomain(data []Point) (result []point.Domain) {
	result = []point.Domain{}
	for _, po := range data {
		result = append(result, po.ToDomain())
	}
	return
}

func FromDomain(domain point.Domain) Point {
	return Point{
		Id:        domain.Id,
		UserId:    domain.UserId,
		Point:     domain.Point,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
