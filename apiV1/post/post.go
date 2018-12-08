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
	p := postmodel.New(user.ID, f.Message, f.IsPrivate)
	if err := p.Save(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	return c.JSON(200, echo.Map{
		"message": "post successfully",
		"id":      p.ID.Hex(),
	})
}

func List(c echo.Context) error {
	posts, err := postmodel.Find(bson.M{})
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	response := []echo.Map{}
	for _, p := range posts {
		response = append(response, miniResponse(p))
	}
	return c.JSON(200, echo.Map{
		"posts": posts,
	})
}
