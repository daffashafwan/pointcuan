package usecase

import (
	"github.com/daffashafwan/pointcuan/business/admin"
	"github.com/daffashafwan/pointcuan/model"
)

type AdminUsecaseImpl struct {
	adminRepo admin.AdminRepo
}

func CreateAdminUsecase(adminRepo admin.AdminRepo) admin.AdminUsecase {
	return &AdminUsecaseImpl{adminRepo}
}

func (e *AdminUsecaseImpl) ReadByUsername(username string)(*model.Admin, error) {
	return e.adminRepo.ReadByUsername(username)
}

