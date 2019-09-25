package v1

import (
	"github.com/globalsign/mgo/bson"
	"github.com/jeyem/feedmeed/src/models/post"
	"github.com/jeyem/feedmeed/src/models/user"
	"github.com/labstack/echo"
)

func newPost(c echo.Context) error {
	f := new(postform)
	if err := c.Bind(f); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	user, err := user.LoadByRequest(c)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	p, err := post.New(user, f.Message)
	if err != nil {
		return c.JSON(500, echo.Map{"error": err.Error()})
	}
	return c.JSON(200, echo.Map{
		"message": "post successfully",
		"id":      p.ID.Hex(),
	})
}

func timeline(c echo.Context) error {
	user, err := user.LoadByRequest(c)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	page := 1
	limit := 100
	timeline := post.LoadTimeline(user.ID, page, limit)

	response := []bson.M{}

	for _, t := range timeline {
		response = append(response, t.V1())
	}

	return c.JSON(200, bson.M{
		"posts": response,
		"page":  page,
		"limit": limit,
	})
}

func selfPosts(c echo.Context) error {
	user, err := user.LoadByRequest(c)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	page := 1
	limit := 100

	posts := post.SelfPosts(user.ID, page, limit)
	response := []bson.M{}
	for _, p := range posts {
		response = append(response, p.V1())
	}
	return c.JSON(200, echo.Map{
		"page":  page,
		"limit": limit,
		"posts": posts,
	})
}
