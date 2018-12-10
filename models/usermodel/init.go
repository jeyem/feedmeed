package usermodel

import (
	"github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
	"github.com/jeyem/feedmeed/app"
)

var (
	a   *app.App
	mdb *bolt.Bucket
)

func Init(application *app.App) {
	a = application
	m, err := a.MDB.CreateBucketIfNotExists([]byte("users"))
	if err != nil {
		logrus.Fatal(err)
	}
	mdb = m

	Connections.gc()
}
