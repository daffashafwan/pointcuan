package jwt

import (
	"time"
	"github.com/golang-jwt/jwt"
)

func JWTHelper(name string, isAdmin bool, role string) string{
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["admin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(role))	
	if err != nil {
		return "Error JWT"
	}
	return t
}