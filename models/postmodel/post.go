package postmodel

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	ID        bson.ObjectId `bson:"_id"`
	Sender    bson.ObjectId `bson:"_sender"`
	Message   string        `bson:"message"`
	Hashes    []string      `bson:"hashes"`
	IsPrivate bool          `bson:"is_private"`
	Created   time.Time     `bson:"created"`
}

func (p *Post) Save() error {
	if p.ID.Valid() {
		return a.DB.Update(p)
	}
	return a.DB.Create(p)
}

func New(sender bson.ObjectId, message string, private bool) *Post {
	p := new(Post)
	p.Sender = sender
	p.IsPrivate = private
	p.Message = message

	return p
}

func Find(filter bson.M) (posts []Post, err error) {
	err = a.DB.Find(filter).Load(&posts)
	return
}
