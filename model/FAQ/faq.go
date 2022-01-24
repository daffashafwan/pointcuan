package FAQ

import (
	"github.com/daffashafwan/pointcuan/business/FAQ"
	"time"
)

type Faq struct {
	Id        int `gorm:"primaryKey"`
	Question  string
	Answer    string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *Faq) ToDomain() FAQ.Domain {
	return FAQ.Domain{
		Id:        user.Id,
		Question:  user.Question,
		Answer:    user.Answer,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToListDomain(data []Faq) (result []FAQ.Domain) {
	result = []FAQ.Domain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}

func FromDomain(domain FAQ.Domain) Faq {
	return Faq{
		Id:        domain.Id,
		Question:  domain.Question,
		Answer:    domain.Answer,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
