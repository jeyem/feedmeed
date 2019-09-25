package v1

import (
	"github.com/jeyem/feedmeed/src/models/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func routes(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	v1.GET("/socket", socket)
	v1.POST("/auth/login", login)
	v1.POST("/auth/register", register)

	r := v1.Group("/")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret should load from config"),
		Claims:     new(user.JwtClaims),
	}))

	u := r.Group("user")
	u.GET("", currentUser)

	f := r.Group("follow")
	f.PUT("/:target", follow)
	f.GET("/ers", followerList)
	f.GET("/ings", followingList)

	t := r.Group("timeline")
	t.GET("", timeline)

	p := r.Group("post")
	p.POST("", newPost)
	p.GET("", selfPosts)
}
