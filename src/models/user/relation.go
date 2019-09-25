package user

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Relation struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Follower  bson.ObjectId `bson:"_follower"`
	Following bson.ObjectId `bson:"_following"`
	Status    string        `bson:"status"`
	Created   time.Time     `bson:"created"`
}

func (*Relation) C() string {
	return "relation"
}

func (r *Relation) Collection() *mgo.Collection {
	return a.DB.C(r.C())
}

func (r *Relation) Insert() error {
	r.Created = time.Now()
	return r.Collection().Insert(r)
}

func (r *Relation) Save() error {
	return r.Insert()
}
