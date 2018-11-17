package app

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/jeyem/feedmeed/app/config"
	"github.com/jeyem/mogo"
	"github.com/labstack/echo"
)

type App struct {
	DB     *mogo.DB
	Config *config.Config
	HTTP   *echo.Echo
}

func New(c config.Config) *App {
	a := new(App)
	a.Config = &c
	a.DB = a.dbConnection()
	a.HTTP = echo.New()
	return a
}

func (a *App) Run() {
	a.HTTP.Static("static", filepath.Join(a.Config.Views, "static"))
	a.HTTP.Logger.Fatal(a.HTTP.Start(fmt.Sprintf(":%d", a.Config.Port)))
}

func (a *App) dbConnection() *mogo.DB {
	uri := fmt.Sprintf("%s:%d/%s",
		a.Config.MongoHost,
		a.Config.MongoPort,
		a.Config.MongoDB)
	db, err := mogo.Conn(uri)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
