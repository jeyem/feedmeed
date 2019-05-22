package postmodel

import (
	"time"

	"github.com/jeyem/feedmeed/models/usermodel"

	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	ID      bson.ObjectId `bson:"_id" json:"_id"`
	Sender  bson.ObjectId `bson:"sender" json:"sender"`
	Message string        `bson:"message" json:"message"`
	Hashes  []string      `bson:"hashes" json:"hashes"`
	Created time.Time     `bson:"created" json:"created"`
}

func (p *Post) Save() error {
	if p.ID.Valid() {
		return a.DB.Update(p)
	}
	p.Created = time.Now()
	return a.DB.Create(p)
}

func New(sender *usermodel.User, message string) (*Post, error) {
	p := new(Post)
	p.Sender = sender.ID
	p.Message = message
	p.Hashes = HashParser(message)
	if err := p.Save(); err != nil {
		return nil, err
	}
	PushTimeline(p, sender)
	return p, nil
}

func SelfPosts(sender bson.ObjectId, page, limit int) (posts []Post) {
	// TODO: fix pagination
	a.DB.Find(bson.M{
		"sender": sender,
	}).Load(&posts)
	return posts
}

func HashParser(msg string) (hashes []string) {
	return hashes
}
