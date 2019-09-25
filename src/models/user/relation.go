package user

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
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
	r.ID = bson.NewObjectId()
	r.Created = time.Now()
	return r.Collection().Insert(r)
}

func (r *Relation) Save() error {
	return r.Insert()
}
