package v1

import (
	"strconv"

	"github.com/globalsign/mgo/bson"

	"github.com/jeyem/feedmeed/src/models/user"
	"github.com/labstack/echo"
)

func follow(c echo.Context) error {
	u, err := user.LoadByRequest(c)
	if err != nil {
		return err
	}
	target := new(user.User)
	if err := target.LoadByUsername(c.Param("target")); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	if err := u.Follow(target); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{"message": "success"})
}

func followerList(c echo.Context) error {
	u, err := user.LoadByRequest(c)
	if err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit > 100 {
		limit = 100
	}

	followers := u.FollowersObjs(page, limit)
	var response = []bson.M{}
	for _, f := range followers {
		response = append(response, f.V1())
	}
	return c.JSON(200, echo.Map{
		"followers": response,
		"page":      page,
		"limit":     limit,
	})
}

func followingList(c echo.Context) error {
	u, err := user.LoadByRequest(c)
	if err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit > 100 {
		limit = 100
	}

	followings := u.FollowingsObjs(page, limit)
	var response = []bson.M{}
	for _, f := range followings {
		response = append(response, f.V1())
	}
	return c.JSON(200, echo.Map{
		"followings": response,
		"page":       page,
		"limit":      limit,
	})
}

func currentUser(c echo.Context) error {
	u, err := user.LoadByRequest(c)
	if err != nil {
		return err
	}
	return c.JSON(200, u.V1())
}
