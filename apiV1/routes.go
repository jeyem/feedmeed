package apiV1

import (
	"github.com/jeyem/feedmeed/apiV1/post"
	"github.com/jeyem/feedmeed/apiV1/user"
	"github.com/jeyem/feedmeed/app"
	"github.com/jeyem/feedmeed/models"
	"github.com/jeyem/feedmeed/models/usermodel"
	"github.com/labstack/echo/middleware"
)

func Register(a *app.App) {

	models.RegisterAllModels(a)

	v1 := a.HTTP.Group("/api/v1")
	v1.GET("/socket", SocketConnect)
	v1.POST("/auth/login", user.Login)
	v1.POST("/auth/register", user.Register)

	r := v1.Group("/")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret should load from config"),
		Claims:     new(usermodel.JwtClaims),
	}))

	u := r.Group("user")
	u.GET("", user.CurrentUser)

	f := r.Group("follow")
	f.PUT("/:target", user.Follow)
	f.GET("/ers", user.FollowerList)
	f.GET("/ings", user.FollowingList)

	t := r.Group("timeline")
	t.GET("", post.Timeline)

	p := r.Group("post")
	p.POST("", post.New)
	p.GET("", post.SelfPosts)

}
