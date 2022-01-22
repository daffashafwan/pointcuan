package requests

import (
	FAQ "github.com/daffashafwan/pointcuan/business/FAQ"
	"strconv"
)


type FaqRequest struct {
	Question string `json:"question"`
	Answer string `json:"answer"`
	Status string `json:"status"`
}

func (cat *FaqRequest) ToDomain() FAQ.Domain {
	stat, _ := strconv.Atoi(cat.Status)
	return FAQ.Domain{
		Question: cat.Question,
		Answer: cat.Answer,
		Status: stat,
	}
}
