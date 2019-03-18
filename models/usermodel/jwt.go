package usermodel

import (
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func LoadByRequest(c echo.Context) (*User, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtClaims)
	u := new(User)
	if err := u.LoadByUsername(claims.Username); err != nil {
		return nil, err
	}
	return u, nil
}

func LoadByToken(tokenStr string) (*User, error) {
	token, _ := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		claims := token.Claims.(*JwtClaims)
		fmt.Println("---------", claims.Username)
		return claims, nil
	})
	u := new(User)
	claims := token.Claims.(*JwtClaims)
	fmt.Println(claims.Username, "---------*****")

	if err := u.LoadByUsername(claims.Username); err != nil {
		return nil, err
	}
	return u, nil
}

func GetToken(c echo.Context) string {
	req := c.Request()
	cleared := strings.Replace(req.Header.Get("Authorization"), " ", "", -1)
	return strings.Replace(cleared, "Bearer", "", -1)
}
