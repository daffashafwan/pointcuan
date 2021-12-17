package handler

import (
	"net/http"
	"strconv"

	"github.com/daffashafwan/pointcuan/business/user"
	"github.com/daffashafwan/pointcuan/helpers/jwt"
	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/daffashafwan/pointcuan/model"
	"github.com/daffashafwan/pointcuan/route"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserUsecase user.UserUsecase
}

func CreateUserHandler(r *echo.Echo, UserUsecase user.UserUsecase) {
	Handler := UserHandler{UserUsecase}
	r.POST("/user/login",Handler.login )
	r.POST("/user/register",Handler.addUser )
	r.GET("/user", Handler.viewUsers, route.IsLoggedInAdmin)
	r.GET("user/:id", Handler.viewUserId)
	r.PUT("/user/:id", Handler.editUser)
	r.DELETE("/user/:id", Handler.deleteUser)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (e *UserHandler) login(c echo.Context) error{
	var user = model.User{}
	err := c.Bind(&user)
	username := user.Username
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "Error Login")
		
	}
	users, err := e.UserUsecase.ReadByUsername(username)
	if err != nil {
		return response.ErrorResponse(c, http.StatusNotFound, err.Error())
		
	}
	if !CheckPasswordHash(user.Password,users.Password) {
		return response.ErrorResponse(c, http.StatusNotFound, "Password Tidak Sesuai")
	}

	return response.SuccessResponse(c, http.StatusOK, jwt.JWTHelper(users.Name, false, "user"))
}

func (e *UserHandler) addUser(c echo.Context) error{
	var user = model.User{}
	err := c.Bind(&user)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, "Oopss server someting wrong")
		
	}
	if user.Id != 0 {
		return response.ErrorResponse(c, http.StatusBadRequest, "input not permitted")
		
	}

	if user.Name == "" || user.Email == "" {
		return response.ErrorResponse(c, http.StatusBadRequest, "column cannot be empty")
		
	}
	newUser, err := e.UserUsecase.Create(&user)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		
	}
	return response.SuccessResponse(c,http.StatusOK, newUser)
}

func (e *UserHandler) viewUsers(c echo.Context)error {
	users, err := e.UserUsecase.ReadAll()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		
	}
	if len(*users) == 0 {
		return response.ErrorResponse(c, http.StatusNotFound, "list user is empty")
		
	}
	return response.SuccessResponse(c,http.StatusOK, users)
}

func (e *UserHandler) viewUserId(c echo.Context) error{
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "id has be number")
		
	}
	user, err := e.UserUsecase.ReadById(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusNotFound, err.Error())
		
	}
	return response.SuccessResponse(c,http.StatusOK ,user)
}

func (e *UserHandler) editUser(c echo.Context)error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "id has be number")
		
	}
	_, err = e.UserUsecase.ReadById(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusNotFound, err.Error())
		
	}
	var tempUser = model.User{}
	err = c.Bind(&tempUser)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, "Oopss server someting wrong")
		
	}
	if tempUser.Id != 0 {
		return response.ErrorResponse(c, http.StatusBadRequest, "input not permitted")
		
	}
	if tempUser.Email == "" || tempUser.Name == "" {
		return response.ErrorResponse(c, http.StatusBadRequest, "column cannot be empty")
		
	}
	user, err := e.UserUsecase.Update(id, &tempUser)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		
	}
	return response.SuccessResponse(c, http.StatusOK,user)
}

func (e *UserHandler) deleteUser(c echo.Context) error{
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "id has be number")
		
	}
	err = e.UserUsecase.Delete(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusNotFound, err.Error())
		
	}
	return response.SuccessResponse(c, http.StatusOK,"success delete data")
}