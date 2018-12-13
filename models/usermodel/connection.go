package usermodel

import (
	"fmt"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"

	"gopkg.in/mgo.v2/bson"
)

var Connections *Sockets

type BroadCaster struct {
	CasterType string `json:"casterType"`
	Data       []byte `json:"data"`
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

func (b *BroadCaster) Message() string {
	return fmt.Sprintf("%s:%s", b.CasterType, string(b.Data))
}

func (s *Sockets) New(c *Socket) {
	s.Lock()
	defer s.Unlock()
	c.lastTime = time.Now().UnixNano()
	s.interfaces[c.Token] = c
	if err := MakeOnline(c.ID); err != nil {
		logrus.Warning(c.ID.Hex(), " status update error")
	}

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
