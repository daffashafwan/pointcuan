package responses

import (
	"time"

	pcrcrud "github.com/daffashafwan/pointcuan/business/pcr_crud"
)

type PcrResponse struct {
	Id        int       `json:"id"`
	NilaiPcr int `json:"nilaiPcr"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain pcrcrud.Domain) PcrResponse {
	return PcrResponse{
		Id:        domain.Id,
		NilaiPcr: domain.NilaiPcr,
		UpdatedAt: domain.UpdatedAt,
	}
}
