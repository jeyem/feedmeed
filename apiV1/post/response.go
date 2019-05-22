package post

import (
	"github.com/jeyem/feedmeed/models/postmodel"
	"github.com/labstack/echo"
)

func miniResponse(p postmodel.Post) echo.Map {
	return echo.Map{
		"id":      p.ID,
		"message": p.Message,
		"hashes":  p.Hashes,
	}
}

func miniResponseTimeline(t postmodel.Timeline) echo.Map {
	return echo.Map{
		"id":      t.ID,
		"post":    miniResponse(t.Post),
		"seen":    t.Seen,
		"created": t.Created,
	}
}
