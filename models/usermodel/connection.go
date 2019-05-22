package usermodel

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"

	"gopkg.in/mgo.v2/bson"
)

var Connections *Sockets

type BroadCaster struct {
	Type string `json:"casterType"`
	Data []byte `json:"data"`
}

type Socket struct {
	ID       bson.ObjectId
	Token    string
	Caster   chan *BroadCaster
	lastTime int64
}

type Sockets struct {
	interfaces map[string]*Socket
	sync.Mutex
}

func CastByID(userID bson.ObjectId, castType string, v interface{}) {
	sessions := GetSessions(userID)
	for _, sess := range sessions {
		if err := CastByToken(sess.Token, castType, v); err != nil {
			logrus.Warn("casting token -->", err)
		}
	}
}

func CastByToken(token, castType string, v interface{}) error {
	socket := Connections.Get(token)
	if socket == nil {
		return errors.New("there is no live connection")
	}
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	c := new(BroadCaster)
	c.Type = castType
	c.Data = data
	socket.Caster <- c
	return nil
}

func (b *BroadCaster) Message() string {
	return fmt.Sprintf("%s:%s", b.Type, string(b.Data))
}

func (s *Sockets) New(c *Socket) {
	s.Lock()
	defer s.Unlock()
	c.lastTime = time.Now().UnixNano()
	s.interfaces[c.Token] = c
}

func (s *Sockets) Get(token string) *Socket {
	s.Lock()
	defer s.Unlock()
	c, ok := s.interfaces[token]
	if !ok {
		return nil
	}
	c.lastTime = time.Now().UnixNano()
	s.interfaces[token] = c
	return c
}

func (s *Sockets) Pull(token string) {
	s.Lock()
	defer s.Unlock()
	delete(s.interfaces, token)
}

func (s *Sockets) gc() {
	go func() {
		for {
			for _, socket := range s.interfaces {
				now := time.Now().UnixNano()
				duration := now - socket.lastTime
				if duration > int64(time.Minute*15) {
					s.Pull(socket.Token)
				}
			}
			time.Sleep(time.Minute * 10)
		}
	}()
}
