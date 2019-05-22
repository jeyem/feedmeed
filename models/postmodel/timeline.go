package postmodel

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jeyem/feedmeed/models/usermodel"

	"gopkg.in/mgo.v2/bson"
)

type Timeline struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Post    Post          `bson:"post"`
	User    bson.ObjectId `bson:"_user"`
	Seen    bool          `bson:"seen"`
	Created time.Time     `bson:"created"`
}

func (t *Timeline) Save() error {
	t.Created = time.Now()
	return a.DB.Create(t)
}

func LoadTimeline(userID bson.ObjectId, page, limit int) (timeline []Timeline) {
	// TODO: fix pagination
	a.DB.Find(bson.M{"_user": userID}).Load(&timeline)
	return timeline
}

func PushTimeline(p *Post, u *usermodel.User) chan bool {
	doneSignal := make(chan bool, 1)
	go func() {
		iter := u.StreamFollowersObjs()
		follower := new(usermodel.User)
		for iter.Next(follower) {
			t := new(Timeline)
			t.User = follower.ID
			t.Post = *p
			if err := t.Save(); err != nil {
				logrus.Error("on pushing timeline -->", err)
			}
			usermodel.CastByID(follower.ID, "timeline", t)
		}
		doneSignal <- true
	}()
	return doneSignal
}
