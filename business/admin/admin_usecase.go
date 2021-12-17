package admin

import "github.com/daffashafwan/pointcuan/model"

type AdminUsecase interface {
	ReadByUsername(username string)(*model.Admin, error)
}