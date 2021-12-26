package pcrcrud

import (
	"time"

	pcrcrud "github.com/daffashafwan/pointcuan/business/pcr_crud"
)

type Pcrcrud struct {
	Id        int `gorm:"primaryKey]"`
	PcrValue int
	UpdatedAt time.Time
}

func (pcr *Pcrcrud) ToDomain() pcrcrud.Domain {
	return pcrcrud.Domain {
		Id:        pcr.Id,
		PcrValue:  pcr.PcrValue,
		UpdatedAt: pcr.UpdatedAt,
	}
}

func FromDomain(domain pcrcrud.Domain) Pcrcrud {
	return Pcrcrud{
		Id:        domain.Id,
		PcrValue: domain.PcrValue,
		UpdatedAt: domain.UpdatedAt,
	}
}
