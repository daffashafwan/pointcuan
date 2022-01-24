package responses

import (
	"strconv"
	"time"

	FAQ "github.com/daffashafwan/pointcuan/business/FAQ"
)

type FaqResponse struct {
	Id        int       `json:"id"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain FAQ.Domain) FaqResponse {
	stat := strconv.Itoa(domain.Status)
	return FaqResponse{
		Id:        domain.Id,
		Question:  domain.Question,
		Answer:    domain.Answer,
		Status:    stat,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(domain []FAQ.Domain) []FaqResponse {
	var list []FaqResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
