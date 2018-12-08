package user

import (
	"github.com/jeyem/feedmeed/models/usermodel"
	"github.com/labstack/echo"
)

func miniResponse(u *usermodel.User) echo.Map {
	return echo.Map{
		"id":          u.ID,
		"username":    u.Username,
		"displayName": u.DisplayName(),
	}
}
