package v1

import (
	"github.com/asaskevich/govalidator"
	"github.com/jeyem/feedmeed/src/models/user"
	"github.com/jeyem/passwd"
	"github.com/labstack/echo"
)

func login(c echo.Context) error {
	form := new(loginForm)
	if err := c.Bind(form); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	if _, err := govalidator.ValidateStruct(form); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	u := new(user.User)
	if err := u.AuthByUsername(form.Identifier, form.Password); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	t, err := u.CreateToken(c)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "login successfully",
		"token":   t,
		"user":    u.V1(),
	})
}

func register(c echo.Context) error {
	form := new(registerForm)
	if err := c.Bind(form); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	if _, err := govalidator.ValidateStruct(form); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	u := new(user.User)
	u.Username = form.Username
	u.Nikname = form.Nikname
	u.Password = passwd.Make(form.Password)
	if err := u.Save(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	t, err := u.CreateToken(c)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	return c.JSON(200, echo.Map{
		"message": "login successfully",
		"token":   t,
		"user":    u.V1(),
	})
}

func check(c echo.Context) error {
	u, err := user.LoadByRequest(c)
	if err != nil {
		return err
	}
	return c.JSON(200, echo.Map{
		"message":  "everything is OK",
		"username": u.Username,
	})
}
