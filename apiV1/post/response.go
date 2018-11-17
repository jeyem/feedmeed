package post

import (
	"github.com/jeyem/feedmeed/models/post"
	"github.com/labstack/echo"
)

func miniResponse(p post.Post) echo.Map {
	return echo.Map{
		"id":      p.ID,
		"message": p.Message,
	}
}
