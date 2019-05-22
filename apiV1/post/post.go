package post

import (
	"github.com/jeyem/feedmeed/models/postmodel"
	"github.com/jeyem/feedmeed/models/usermodel"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func New(c echo.Context) error {
	f := new(form)
	if err := c.Bind(f); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	user, err := usermodel.LoadByRequest(c)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	p, err := postmodel.New(user, f.Message)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	return c.JSON(200, echo.Map{
		"message": "post successfully",
		"id":      p.ID.Hex(),
	})
}

func Timeline(c echo.Context) error {
	user, err := usermodel.LoadByRequest(c)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	page := 1
	limit := 100
	timeline := postmodel.LoadTimeline(user.ID, page, limit)

	response := []echo.Map{}

	for _, item := range timeline {
		response = append(response, miniResponseTimeline(item))
	}

	return c.JSON(200, bson.M{
		"posts": response,
		"page":  page,
		"limit": limit,
	})
}

func SelfPosts(c echo.Context) error {
	user, err := usermodel.LoadByRequest(c)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	page := 1
	limit := 100

	posts := postmodel.SelfPosts(user.ID, page, limit)
	response := []echo.Map{}
	for _, p := range posts {
		response = append(response, miniResponse(p))
	}
	return c.JSON(200, echo.Map{
		"page":  page,
		"limit": limit,
		"posts": posts,
	})
}
