package handler

import (
	"net/http"
	"time"
	"strconv"
	"golang.org/x/crypto/bcrypt"
	"github.com/daffashafwan/pointcuan/helpers"
	"github.com/daffashafwan/pointcuan/model"
	"github.com/daffashafwan/pointcuan/user"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
)

type UserHandler struct {
	UserUsecase user.UserUsecase
}

func CreateUserHandler(r *echo.Echo, UserUsecase user.UserUsecase) {
	Handler := UserHandler{UserUsecase}
	r.POST("/user/login",Handler.login )
	r.POST("/user",Handler.addUser )
	r.GET("/user", Handler.viewUsers)
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
		return helpers.ErrorResponse(c, http.StatusBadRequest, "Error Login")
		
	}
	users, err := e.UserUsecase.ReadByUsername(username)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusNotFound, err.Error())
		
	}
	if !CheckPasswordHash(user.Password,users.Password) {
		return helpers.ErrorResponse(c, http.StatusNotFound, "Password Tidak Sesuai")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = users.Name
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("user"))	
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Access Token": t,
	})
}

func (e *UserHandler) addUser(c echo.Context) error{
	var user = model.User{}
	err := c.Bind(&user)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Oopss server someting wrong")
		
	}
	if user.Id != 0 {
		return helpers.ErrorResponse(c, http.StatusBadRequest, "input not permitted")
		
	}

	if user.Name == "" || user.Email == "" {
		return helpers.ErrorResponse(c, http.StatusBadRequest, "column cannot be empty")
		
	}
	newUser, err := e.UserUsecase.Create(&user)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		
	}
	return helpers.SuccessResponse(c,http.StatusOK, newUser)
}

func (e *UserHandler) viewUsers(c echo.Context)error {
	users, err := e.UserUsecase.ReadAll()
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		
	}
	if len(*users) == 0 {
		return helpers.ErrorResponse(c, http.StatusNotFound, "list user is empty")
		
	}
	return helpers.SuccessResponse(c,http.StatusOK, users)
}

func (e *UserHandler) viewUserId(c echo.Context) error{
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, "id has be number")
		
	}
	user, err := e.UserUsecase.ReadById(id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusNotFound, err.Error())
		
	}
	return helpers.SuccessResponse(c,http.StatusOK ,user)
}

func (e *UserHandler) editUser(c echo.Context)error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, "id has be number")
		
	}
	_, err = e.UserUsecase.ReadById(id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusNotFound, err.Error())
		
	}
	var tempUser = model.User{}
	err = c.Bind(&tempUser)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Oopss server someting wrong")
		
	}
	if tempUser.Id != 0 {
		return helpers.ErrorResponse(c, http.StatusBadRequest, "input not permitted")
		
	}
	if tempUser.Email == "" || tempUser.Name == "" {
		return helpers.ErrorResponse(c, http.StatusBadRequest, "column cannot be empty")
		
	}
	user, err := e.UserUsecase.Update(id, &tempUser)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		
	}
	return helpers.SuccessResponse(c, http.StatusOK,user)
}

func (e *UserHandler) deleteUser(c echo.Context) error{
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, "id has be number")
		
	}
	err = e.UserUsecase.Delete(id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusNotFound, err.Error())
		
	}
	return helpers.SuccessResponse(c, http.StatusOK,"success delete data")
}