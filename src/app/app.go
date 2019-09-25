package app

import (
	"fmt"
	"path/filepath"

	"github.com/boltdb/bolt"
	"github.com/globalsign/mgo"
	"github.com/jeyem/feedmeed/src/app/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

type App struct {
	DB     *mgo.Database
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

func (a *App) dbConnection() *mgo.Database {
	info := &mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s:%d", a.Config.MongoHost, a.Config.MongoPort)},
		Database: a.Config.MongoDB,
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		logrus.Fatal("mongo db ", err)
	}
	return session.DB(info.Database)
}
