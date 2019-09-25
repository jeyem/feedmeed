package user

import (
	"errors"
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
		return claims, nil
	})
	u := new(User)
	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New("converting token failed")
	}

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
