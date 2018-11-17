package main

import (
	"github.com/jeyem/feedmeed/apiV1"
	"github.com/jeyem/feedmeed/app"
	"github.com/jeyem/feedmeed/app/config"
)

func main() {
	a := app.New(config.DefaultConfig)
	apiV1.Register(a)
	a.Run()
}
