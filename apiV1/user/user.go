package user

import (
	"github.com/jeyem/feedmeed/models/usermodel"
	"github.com/labstack/echo"
)

func CurrentUser(c echo.Context) error {
	u, err := usermodel.LoadByRequest(c)
	if err != nil {
		return err
	}
	return c.JSON(200, echo.Map{
		"user": miniResponse(u),
	})
}
