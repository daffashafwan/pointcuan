package repo

import (
	"fmt"

	"github.com/daffashafwan/pointcuan/business/admin"
	"github.com/daffashafwan/pointcuan/model"

	"github.com/jinzhu/gorm"
)

type AdminRepoImpl struct {
	DB *gorm.DB
}

func CreateAdminRepo(DB *gorm.DB) admin.AdminRepo {
	return &AdminRepoImpl{DB}
}

func (e *AdminRepoImpl) ReadByUsername(username string)(*model.Admin, error) {
	var admin = model.Admin{}
	err := e.DB.Table("admins").Where("username = ?", username).First(&admin).Error
	if err != nil {
		fmt.Printf("[AdminRepoImpl.ReadById] error execute query %v \n", err)
		return nil, fmt.Errorf("username is not exsis")
	}
	return &admin, nil
}