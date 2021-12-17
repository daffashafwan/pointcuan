package admin

import "github.com/daffashafwan/pointcuan/model"

type AdminRepo interface {
	ReadByUsername(username string)(*model.Admin, error)
}