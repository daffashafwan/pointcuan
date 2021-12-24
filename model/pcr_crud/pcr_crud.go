package pcrcrud

import (
	"time"

	pcrcrud "github.com/daffashafwan/pointcuan/business/pcr_crud"
)

type Pcrcrud struct {
	Id        int `gorm:"primaryKey]"`
	NilaiPcr int
	UpdatedAt time.Time
}

func (pcr *Pcrcrud) ToDomain() pcrcrud.Domain {
	return pcrcrud.Domain {
		Id:        pcr.Id,
		NilaiPcr:  pcr.NilaiPcr,
		UpdatedAt: pcr.UpdatedAt,
	}
}

func FromDomain(domain pcrcrud.Domain) Pcrcrud {
	return Pcrcrud{
		Id:        domain.Id,
		NilaiPcr: domain.NilaiPcr,
		UpdatedAt: domain.UpdatedAt,
	}
}
