package post

import (
	"github.com/jeyem/feedmeed/models/postmodel"
	"github.com/labstack/echo"
)

func miniResponse(p postmodel.Post) echo.Map {
	return echo.Map{
		"id":      p.ID,
		"message": p.Message,
	}
}
