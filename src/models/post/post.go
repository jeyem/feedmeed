package post

import (
	"time"

	"github.com/jeyem/feedmeed/src/models/user"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	ID      bson.ObjectId `bson:"_id" json:"_id"`
	Sender  bson.ObjectId `bson:"sender" json:"sender"`
	Message string        `bson:"message" json:"message"`
	Tags    []string      `bson:"tags" json:"tags"`
	Created time.Time     `bson:"created" json:"created"`
}

func (*Post) C() string {
	return "post"
}

func (p *Post) Collection() *mgo.Collection {
	return a.DB.C(p.C())
}

func (p *Post) Insert() error {
	p.Created = time.Now()
	return p.Collection().Insert(p)
}

func (p *Post) Save() error {
	return p.Insert()
}

func (p *Post) V1() bson.M {
	return bson.M{
		"id":      p.ID.Hex(),
		"sender":  p.Sender.Hex(),
		"message": p.Message,
		"tags":    p.Tags,
	}
}

func New(sender *user.User, message string) (*Post, error) {
	p := new(Post)
	p.Sender = sender.ID
	p.Message = message
	p.Tags = TagsParser(message)
	if err := p.Save(); err != nil {
		return nil, err
	}
	PushTimeline(p, sender)
	return p, nil
}

func SelfPosts(sender bson.ObjectId, page, limit int) (posts []Post) {
	// TODO: fix pagination
	p := new(Post)
	p.Collection().Find(bson.M{"sender": sender}).All(posts)
	return posts
}

func TagsParser(msg string) (tags []string) {
	return
}
