package main

import (
	"github.com/jeyem/feedmeed/src/api/v1"
	"github.com/jeyem/feedmeed/src/app"
	"github.com/jeyem/feedmeed/src/app/config"
)

func main() {
	a := app.New(config.DefaultConfig)
	v1.Register(a)
	a.Run()
}
