package app

import (
	"fmt"
	"path/filepath"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/Sirupsen/logrus"
	"github.com/jeyem/feedmeed/app/config"
	"github.com/jeyem/mogo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	if a.Config.Debug {
		a.HTTP.Use(middleware.CORS())
	}
	return a
}

func (a *App) Run() {
	a.HTTP.Static("static", filepath.Join(a.Config.Views, "static"))
	a.HTTP.Logger.Fatal(a.HTTP.Start(fmt.Sprintf(":%d", a.Config.Port)))
}

func (a *App) dbConnection() *mogo.DB {
	db, err := mogo.Conn(&mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s:%d", a.Config.MongoHost, a.Config.MongoPort)},
		Timeout:  60 * time.Second,
		Database: a.Config.MongoDB,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	return db
}
