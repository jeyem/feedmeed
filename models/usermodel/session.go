package usermodel

import (
	"time"

	"github.com/labstack/echo"

	"gopkg.in/mgo.v2/bson"
)

type Session struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	ClientInfo string        `bson:"device"`
	Token      string        `bson:"token"`
	User       bson.ObjectId `bson:"_user"`
	Created    time.Time     `bson:"created"`
}

func (s *Session) Save() error {
	s.Created = time.Now()
	return a.DB.Create(s)
}

func GetSessions(userID bson.ObjectId) (sessions []Session) {
	a.DB.Find(bson.M{"_user": userID}).Load(&sessions)
	return sessions
}

func CreateSession(c echo.Context, u *User, token string) error {
	sess := new(Session)
	sess.User = u.ID
	sess.Token = token
	return sess.Save()
}
