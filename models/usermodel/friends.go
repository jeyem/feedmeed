package usermodel

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Friend struct {
	ID            bson.ObjectId `bson:"_id"`
	Requester     bson.ObjectId `bson:"requester"`
	PendingOnUser bson.ObjectId `bson:"pending_on_user"`
	Accepted      bool          `bson:"accepted"`
	Created       time.Time     `bson:"created"`
}

func (f *Friend) Save() error {
	if f.ID.Valid() {
		return a.DB.Update(f)
	}
	f.Created = time.Now()
	return a.DB.Create(f)
}
