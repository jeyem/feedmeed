package postmodel

import (
	"time"

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

func New(sender bson.ObjectId, message string) (*Post, error) {
	p := new(Post)
	p.Sender = sender
	p.Message = message
	p.Hashes = HashParser(message)
	if err := p.Save(); err != nil {
		return nil, err
	}
	// WriteTimeLines(sender, p)
	return p, nil
}

func SelfPosts(sender bson.ObjectId, page, limit int) (posts []Post) {
	a.DB.Find(bson.M{
		"sender": sender,
	}).Load(&posts)
	return posts
}

func HashParser(msg string) (hashes []string) {
	return hashes
}
