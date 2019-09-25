package post

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/jeyem/feedmeed/src/models/user"
	"github.com/sirupsen/logrus"
)

type Timeline struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Post    Post          `bson:"post"`
	User    bson.ObjectId `bson:"_user"`
	Seen    bool          `bson:"seen"`
	Created time.Time     `bson:"created"`
}

func (*Timeline) C() string {
	return "timeline"
}

func (t *Timeline) Collection() *mgo.Collection {
	return a.DB.C(t.C())
}

func (t *Timeline) Insert() error {
	t.ID = bson.NewObjectId()
	t.Created = time.Now()
	return t.Collection().Insert(t)
}

func (t *Timeline) Save() error {
	return t.Insert()
}

func (t *Timeline) V1() bson.M {
	return bson.M{
		"id":      t.ID,
		"post":    t.Post.V1(),
		"seen":    t.Seen,
		"created": t.Created,
	}
}

func LoadTimeline(userID bson.ObjectId, page, limit int) (timeline []Timeline) {
	// TODO: fix pagination
	// skip := (page - 1) * limit
	t := new(Timeline)
	t.Collection().Find(bson.M{"_user": userID}).All(&timeline)
	return timeline
}

func PushTimeline(p *Post, u *user.User) chan bool {
	doneSignal := make(chan bool, 1)
	go func() {
		iter := u.StreamFollowersObjs()
		relation := new(user.Relation)
		for iter.Next(relation) {
			t := new(Timeline)
			t.User = relation.Follower
			t.Post = *p
			if err := t.Save(); err != nil {
				logrus.Error("on pushing timeline -->", err)
			}
			user.CastByID(relation.Follower, "timeline", t)
		}
		doneSignal <- true
	}()
	return doneSignal
}
