package user

import (
	"time"

	"github.com/labstack/echo"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Session struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	ClientInfo string        `bson:"device"`
	Token      string        `bson:"token"`
	User       bson.ObjectId `bson:"_user"`
	Created    time.Time     `bson:"created"`
}

func (*Session) C() string {
	return "session"
}

func (s *Session) Collection() *mgo.Collection {
	return a.DB.C(s.C())
}

func (s *Session) Insert() error {
	s.ID = bson.NewObjectId()
	s.Created = time.Now()
	return s.Collection().Insert(s)
}

func (s *Session) Save() error {
	return s.Insert()
}

func GetSessions(userID bson.ObjectId) (sessions []Session) {
	s := new(Session)
	s.Collection().Find(bson.M{"_user": userID}).All(&sessions)
	return sessions
}

func CreateSession(c echo.Context, u *User, token string) error {
	sess := new(Session)
	sess.User = u.ID
	sess.Token = token
	return sess.Save()
}
