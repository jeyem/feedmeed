package user

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/jeyem/passwd"
	"github.com/labstack/echo"
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

func (*User) C() string {
	return "user"
}

func (u *User) Collection() *mgo.Collection {
	return a.DB.C(u.C())
}

func (u *User) Insert() error {
	u.Created = time.Now()
	return u.Collection().Insert(u)
}

func (u *User) Save() error {
	return u.Insert()
}

func (u *User) V1() bson.M {
	return bson.M{
		"id":          u.ID.Hex(),
		"username":    u.Username,
		"displayName": u.DisplayName(),
		"followers":   u.Followers,
		"followings":  u.Followings,
	}
}

func (u *User) DisplayName() string {
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

func (u *User) LoadByUsername(username string) error {
	return u.Collection().Find(bson.M{"username": username}).One(u)
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
	if err := CreateSession(c, u, t); err != nil {
		return t, err
	}
	return t, nil
}

func (u *User) Follow(target *User) error {
	target.Followers++
	if err := target.Save(); err != nil {
		return err
	}
	u.Followings++
	if err := u.Save(); err != nil {
		return err
	}
	r := Relation{
		Follower:  u.ID,
		Following: target.ID,
		Status:    follow,
	}
	return r.Save()
}

func (u *User) StreamFollowersObjs() *mgo.Iter {
	r := new(Relation)
	return r.Collection().Find(bson.M{
		"_following": u.ID,
		"status":     follow,
	}).Iter()
}

func (u *User) FollowersObjs(page, limit int) (users []User) {
	relations := []Relation{}
	r := new(Relation)
	r.Collection().Find(bson.M{
		"_following": u.ID,
		"status":     follow,
	}).All(&relations)
	var ids []bson.ObjectId
	for _, r := range relations {
		ids = append(ids, r.Follower)
	}
	u.Collection().Find(bson.M{"_id": bson.M{"$in": ids}}).All(&users)
	return users
}

func (u *User) FollowingsObjs(page, limit int) (users []User) {
	relations := []Relation{}
	r := new(Relation)
	r.Collection().Find(bson.M{
		"_follower": u.ID,
		"status":    follow,
	}).All(&relations)
	var ids []bson.ObjectId
	for _, r := range relations {
		ids = append(ids, r.Following)
	}
	u.Collection().Find(bson.M{"_id": bson.M{"$in": ids}}).All(&users)
	return users
}

func Load(username string) (*User, error) {
	u := new(User)
	if err := u.LoadByUsername(username); err != nil {
		return nil, err
	}
	return u, nil
}

func Query(q bson.M) (users []User) {
	u := new(User)
	u.Collection().Find(q).All(&users)
	return users
}

func Search(query string) (users []User) {
	u := new(User)
	u.Collection().Find(bson.M{"$text": bson.M{"$search": query}}).All(&users)
	return users
}
