package usermodel

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jeyem/passwd"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId   `bson:"_id"`
	Username string          `bson:"username"`
	Nikname  string          `bson:"nikname"`
	Password string          `bson:"password"`
	Friends  []bson.ObjectId `bson:"friends"`
}

func (u *User) Save() error {
	if u.ID.Valid() {
		return a.DB.Update(u)
	}
	return a.DB.Create(u)
}

func (u *User) LoadByUsername(username string) error {
	return a.DB.Find(bson.M{"username": username}).Load(u)
}

func (u *User) AuthByUsername(username, password string) error {
	autherr := errors.New("username or password not matched")
	if err := u.LoadByUsername(username); err != nil {
		return autherr
	}
	if ok := passwd.Check(password, u.Password); !ok {
		return autherr
	}
	return nil
}

func (u *User) CreateToken() (string, error) {

	claims := new(JwtClaims)
	claims.Username = u.Username
	claims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret should load from config"))
}
