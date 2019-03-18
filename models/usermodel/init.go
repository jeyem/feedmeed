package usermodel

import (
	"github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
	"github.com/jeyem/feedmeed/app"
)

var (
	a           *app.App
	userBucket  *bolt.Bucket
	tokenBucket *bolt.Bucket
)

func Init(application *app.App) {
	a = application

	// users cache
	ub, err := a.MDB.CreateBucketIfNotExists([]byte("users"))
	if err != nil {
		logrus.Fatal(err)
	}
	userBucket = ub

	// tokens cache
	tb, err := a.MDB.CreateBucketIfNotExists([]byte("tokens"))
	if err != nil {
		logrus.Fatal(err)
	}
	tokenBucket = tb

	Connections = new(Sockets)
	Connections.interfaces = map[string]*Socket{}

	// sockets garbage collector
	Connections.gc()
	a.DB.LoadIndexes(&User{})
}
