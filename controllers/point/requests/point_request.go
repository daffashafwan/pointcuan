package requests

import "github.com/daffashafwan/pointcuan/business/point"

type PointRequest struct {
	UserId int
	Point  float64
}

func (pr *PointRequest) ToDomain() point.Domain {
	return point.Domain{
		UserId: pr.UserId,
		Point:  pr.Point,
	}
}
