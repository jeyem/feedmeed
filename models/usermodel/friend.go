package usermodel

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type FriendRequest struct {
	ID            bson.ObjectId `bson:"_id"`
	Requester     bson.ObjectId `bson:"requester"`
	PendingOnUser bson.ObjectId `bson:"pending_on_user"`
	Accepted      bool          `bson:"accepted"`
	Rejected      bool          `bson:"rejected"`
	Created       time.Time     `bson:"created"`
}

func (f *FriendRequest) Save() error {
	if f.ID.Valid() {
		return a.DB.Update(f)
	}
	f.Created = time.Now()
	return a.DB.Create(f)
}

func LoadPendingFriendRequest(target, requester bson.ObjectId) (*FriendRequest, error) {
	f := new(FriendRequest)
	if err := a.DB.Find(bson.M{
		"requester":       requester,
		"pending_on_user": target,
		"accepted":        false,
		"rejected":        false,
	}).Load(f); err != nil {
		return nil, err
	}
	return f, nil
}
