package usermodel

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Relation struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Follower  bson.ObjectId `bson:"_follower"`
	Following bson.ObjectId `bson:"_following"`
	Status    string        `bson:"status"`
	Created   time.Time     `bson:"created"`
}

func (r *Relation) Save() error {
	return a.DB.Create(r)
}
