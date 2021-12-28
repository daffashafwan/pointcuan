package requests

import pcrcrud "github.com/daffashafwan/pointcuan/business/pcr_crud"

type PcrUpdate struct {
	PcrValue    int `json:"pcrValue"`
}

func (pcr *PcrUpdate) ToDomain() pcrcrud.Domain {
	return pcrcrud.Domain{
		PcrValue:    pcr.PcrValue,
	}
}