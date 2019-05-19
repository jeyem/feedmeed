package usermodel

import (
	"github.com/jeyem/feedmeed/app"
)

var (
	a *app.App
)

func Init(application *app.App) {
	a = application

	Connections = new(Sockets)
	Connections.interfaces = map[string]*Socket{}

	// sockets garbage collector
	Connections.gc()
	a.DB.LoadIndexes(&User{})
}
