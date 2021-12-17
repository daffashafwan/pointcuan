package handler

import (
	"net/http"

	"github.com/daffashafwan/pointcuan/business/admin"
	"github.com/daffashafwan/pointcuan/helpers/jwt"
	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/daffashafwan/pointcuan/model"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type AdminHandler struct {
	AdminUsecase admin.AdminUsecase
}

func CreateAdminHandler(r *echo.Echo, AdminUsecase admin.AdminUsecase) {
	Handler := AdminHandler{AdminUsecase}
	r.POST("/admin/login",Handler.login )
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (e *AdminHandler) login(c echo.Context) error{
	var admin = model.Admin{}
	err := c.Bind(&admin)
	username := admin.Username
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "Error Login")
		
	}
	users, err := e.AdminUsecase.ReadByUsername(username)
	if err != nil {
		return response.ErrorResponse(c, http.StatusNotFound, err.Error())
		
	}
	if !CheckPasswordHash(admin.Password,users.Password) {
		return response.ErrorResponse(c, http.StatusNotFound, "Password Tidak Sesuai")
	}

	return response.SuccessResponse(c, http.StatusOK, jwt.JWTHelper(admin.Name, true, "admin"))
}