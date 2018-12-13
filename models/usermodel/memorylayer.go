package usermodel

import (
	"encoding/json"
	"errors"

	"github.com/Sirupsen/logrus"

	"gopkg.in/mgo.v2/bson"
)

type CacheLayer struct {
	User struct {
		Username    string `json:"username"`
		Nikname     string `json:"nikname"`
		DisplayName string `json:"displayName"`
	} `json:"user"`
	Sessions []string `json:"sessions"`
	IsOnline bool     `json:"is_online"`
	Status   string   `json:"status"`
}

func NewSession(u *User, token string) error {
	key := []byte(u.ID.Hex())
	c := new(CacheLayer)
	data := mdb.Get(key)
	if err := mdb.Put([]byte(token), []byte(u.ID.Hex())); err != nil {
		logrus.Warning("token catch failed")
	}

	if data != nil {
		if err := json.Unmarshal(data, c); err != nil {
			return err
		}
		c.User.Nikname = u.Nikname
		c.User.DisplayName = u.DisplayName()
		c.Sessions = append(c.Sessions, token)
		data, err := json.Marshal(c)
		if err != nil {
			return err
		}
		if err := mdb.Put(key, data); err != nil {
			return err
		}
		return nil
	}

	c.User.Username = u.Username
	c.User.Nikname = u.Nikname
	c.Sessions = []string{token}
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return mdb.Put(key, data)
}

func UpdateStatus(userID bson.ObjectId, status string) error {
	key := []byte(userID.Hex())
	c := new(CacheLayer)
	data := mdb.Get(key)
	if err := json.Unmarshal(data, c); err != nil {
		return err
	}
	c.Status = status

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return mdb.Put(key, data)
}

func GetUsernameFromToken(token string) (string, error) {
	cache := mdb.Get([]byte(token))
	if cache == nil {
		return "", errors.New("not found")
	}
	return string(cache), nil
}

func MakeOnline(userID bson.ObjectId) error {
	key := []byte(userID.Hex())
	c := new(CacheLayer)
	data := mdb.Get(key)
	if err := json.Unmarshal(data, c); err != nil {
		return err
	}
	c.IsOnline = true

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return mdb.Put(key, data)
}
