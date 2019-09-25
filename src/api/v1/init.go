package v1

import (
	"github.com/jeyem/feedmeed/src/app"
	"github.com/jeyem/feedmeed/src/models/post"
	"github.com/jeyem/feedmeed/src/models/user"
)

func Register(a *app.App) {

	post.Init(a)
	user.Init(a)

	routes(a.HTTP)
}
