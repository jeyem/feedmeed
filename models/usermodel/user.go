package usermodel

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jeyem/passwd"
	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID          bson.ObjectId `bson:"_id" json:"_id"`
	Username    string        `bson:"username" json:"username"`
	Nikname     string        `bson:"nikname" json:"nikname"`
	Password    string        `bson:"password" json:"-"`
	ForceStatus bool          `bson:"force_status" json:"force_status"`
	Followers   int           `bson:"followers"`
	Followings  int           `bson:"followings"`
	Created     time.Time     `bson:"created" json:"created"`
}

func (u User) DisplayName() string {
	if u.Nikname != "" {
		return u.Nikname
	}
	return u.Username
}

func (u *User) Meta() []mgo.Index {
	return []mgo.Index{
		{Key: []string{"username"}, Unique: true},
	}
}

func (u *User) Save() error {
	if u.ID.Valid() {
		return a.DB.Update(u)
	}
	u.Created = time.Now()
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

func (u *User) CreateToken(c echo.Context) (string, error) {

	claims := new(JwtClaims)
	claims.Username = u.Username
	claims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret should load from config"))
	if err != nil {
		return t, err
	}
	if err := CreateSession(c, u.ID, t); err != nil {
		return t, err
	}
	return t, nil
}

// func (u *User) Follow(target *User) error {
// 	target.Followers++
// 	if err := target.Save(); err != nil {
// 		return err
// 	}
// 	u.Followings++
// 	if err := u.Save(); err != nil {
// 		return err
// 	}
// 	r := Relation{
// 		Follower:  u.ID,
// 		Following: target.ID,
// 		Status:    follow,
// 	}
// 	return r.Save()
// }

// func (u *User) FollowersObjs(page, limit int) (users []User) {
// 	relations := []Relation{}
// 	a.DB.Find(bson.M{
// 		"following": u.ID,
// 		"status":    follow,
// 	}).Load(&relations)
// 	var ids []bson.ObjectId
// 	for _, r := range relations {
// 		ids = append(ids, r.Follower)
// 	}
// 	a.DB.Find(bson.M{"_id": bson.M{"$in": ids}}).Load(&users)
// 	return users
// }

// func (u *User) FollowingsObjs(page, limit int) (users []User) {
// 	relations := []Relation{}
// 	a.DB.Find(bson.M{
// 		"follower": u.ID,
// 		"status":   follow,
// 	}).Load(&relations)
// 	var ids []bson.ObjectId
// 	for _, r := range relations {
// 		ids = append(ids, r.Following)
// 	}
// 	a.DB.Find(bson.M{"_id": bson.M{"$in": ids}}).Load(&users)
// 	return users
// }

func Load(username string) (*User, error) {
	u := new(User)
	if err := u.LoadByUsername(username); err != nil {
		return nil, err
	}
	return u, nil
}

func Query(q bson.M) (users []User) {
	a.DB.Find(q).Load(&users)
	return users
}

func Search(query string) (users []User) {
	a.DB.Find(bson.M{"$text": bson.M{"$search": query}}).Load(&users)
	return users
}
