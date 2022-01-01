package responses

import (
	"time"

	pcrcrud "github.com/daffashafwan/pointcuan/business/pcr_crud"
)

type PcrResponse struct {
	Id        int       `json:"id"`
	PcrValue float64 `json:"pcrValue"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain pcrcrud.Domain) PcrResponse {
	return PcrResponse{
		Id:        domain.Id,
		PcrValue: domain.PcrValue,
		UpdatedAt: domain.UpdatedAt,
	}
}
