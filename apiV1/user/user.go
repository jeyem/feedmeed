package user

import (
	"strconv"

	"github.com/jeyem/feedmeed/models/usermodel"
	"github.com/labstack/echo"
)

func Follow(c echo.Context) error {
	u, err := usermodel.LoadByRequest(c)
	if err != nil {
		return err
	}
	target := new(usermodel.User)
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

func FollowerList(c echo.Context) error {
	u, err := usermodel.LoadByRequest(c)
	if err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit > 100 {
		limit = 100
	}

	followers := u.FollowersObjs(page, limit)
	var response []echo.Map
	for _, f := range followers {
		response = append(response, miniResponse(&f))
	}
	return c.JSON(200, echo.Map{
		"followers": response,
		"page":      page,
		"limit":     limit,
	})
}

func FollowingList(c echo.Context) error {
	u, err := usermodel.LoadByRequest(c)
	if err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit > 100 {
		limit = 100
	}

	followings := u.FollowingsObjs(page, limit)
	var response []echo.Map
	for _, f := range followings {
		response = append(response, miniResponse(&f))
	}
	return c.JSON(200, echo.Map{
		"followings": response,
		"page":       page,
		"limit":      limit,
	})
}

func CurrentUser(c echo.Context) error {
	u, err := usermodel.LoadByRequest(c)
	if err != nil {
		return err
	}
	return c.JSON(200, echo.Map{
		"user": miniResponse(u),
	})
}
