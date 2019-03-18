package apiV1

import (
	"github.com/jeyem/feedmeed/apiV1/post"
	"github.com/jeyem/feedmeed/apiV1/search"
	"github.com/jeyem/feedmeed/apiV1/user"
	"github.com/jeyem/feedmeed/app"
	"github.com/jeyem/feedmeed/models"
	"github.com/jeyem/feedmeed/models/usermodel"
	"github.com/labstack/echo/middleware"
)

func Register(a *app.App) {

	models.RegisterAllModels(a)

	v1 := a.HTTP.Group("/api/v1")

	v1.POST("/auth/login", user.Login)
	v1.POST("/auth/register", user.Register)

	r := v1.Group("/r")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret should load from config"),
		Claims:     new(usermodel.JwtClaims),
	}))

	v1.GET("/socket", user.SocketConnect)

	u := r.Group("/user")
	u.GET("", user.CurrentUser)

	f := r.Group("/follow")
	f.PUT("/:target", user.FollowRequest)
	f.GET("/ers", user.FollowersList)
	f.GET("/ings", user.FollowingsList)

	t := r.Group("/timeline")
	t.GET("", post.Timeline)

	p := r.Group("/post")
	p.POST("/new", post.New)
	p.GET("/self", post.SelfPosts)

	s := r.Group("/search")
	s.GET("", search.Search)

}
