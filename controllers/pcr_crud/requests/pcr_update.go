package requests

import pcrcrud "github.com/daffashafwan/pointcuan/business/pcr_crud"

type PcrUpdate struct {
	NilaiPcr    int `json:"nilaiPcr"`
}

func (pcr *PcrUpdate) ToDomain() pcrcrud.Domain {
	return pcrcrud.Domain{
		NilaiPcr:    pcr.NilaiPcr,
	}
}
