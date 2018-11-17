package usermodel

import (
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
