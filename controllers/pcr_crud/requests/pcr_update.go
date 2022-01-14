package requests

import (
	pcrcrud "github.com/daffashafwan/pointcuan/business/pcr_crud"
	"strconv"
)
type PcrUpdate struct {
	PcrValue    string `json:"pcrValue"`
}

func (pcr *PcrUpdate) ToDomain() pcrcrud.Domain {
	conv,_:= strconv.ParseFloat(pcr.PcrValue, 64)
	return pcrcrud.Domain{
		PcrValue:  conv ,
	}
}
