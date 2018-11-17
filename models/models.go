package models

import (
	"github.com/jeyem/feedmeed/app"
	"github.com/jeyem/feedmeed/models/postmodel"
	"github.com/jeyem/feedmeed/models/usermodel"
)

func RegisterAllModels(a *app.App) {
	postmodel.Init(a)
	usermodel.Init(a)
}
