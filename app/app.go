package app

import (
	"fmt"
	"path/filepath"

	mgo "gopkg.in/mgo.v2"

	"github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
	"github.com/jeyem/feedmeed/app/config"
	"github.com/jeyem/mogo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type App struct {
	DB     *mogo.DB
	MDB    *bolt.Tx
	Config *config.Config
	HTTP   *echo.Echo
}

func New(c config.Config) *App {
	a := new(App)
	a.Config = &c
	a.DB = a.dbConnection()
	a.MDB = a.memorydbConnection()
	a.HTTP = echo.New()
	if a.Config.Debug {
		a.HTTP.Use(middleware.CORS())
	}
	return a
}

func (a *App) Run() {
	a.HTTP.Use(middleware.Logger())
	a.HTTP.Static("static", filepath.Join(a.Config.Views, "static"))
	a.HTTP.Logger.Fatal(a.HTTP.Start(fmt.Sprintf(":%d", a.Config.Port)))
}

func (a *App) memorydbConnection() *bolt.Tx {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		logrus.Fatal(err)
	}
	tx, err := db.Begin(true)
	if err != nil {
		logrus.Fatal(err)
	}
	return tx
}

func (a *App) dbConnection() *mogo.DB {
	db, err := mogo.Conn(&mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s:%d", a.Config.MongoHost, a.Config.MongoPort)},
		Database: a.Config.MongoDB,
	})

	if err != nil {
		logrus.Fatal(err)
	}
	return db
}
