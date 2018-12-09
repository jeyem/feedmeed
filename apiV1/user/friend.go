package user

import (
	"github.com/jeyem/feedmeed/models/usermodel"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func AddFriendRequest(c echo.Context) error {
	u, err := usermodel.LoadByRequest(c)
	if err != nil {
		return err
	}
	target, err := usermodel.Load(c.Param("target"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	friend := new(usermodel.FriendRequest)
	friend.Requester = u.ID
	friend.PendingOnUser = target.ID
	if err := friend.Save(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	return c.JSON(200, echo.Map{"message": "friend request pending successfully"})
}

func AcceptFriendRequest(c echo.Context) error {
	u, err := usermodel.LoadByRequest(c)
	if err != nil {
		return err
	}
	requester, err := usermodel.Load(c.Param("requester"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	friendReq, err := usermodel.LoadPendingFriendRequest(u.ID, requester.ID)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	if err := u.AddFriend(friendReq); err != nil {
		return err
	}
	return c.JSON(200, echo.Map{"message": "friend acccepted successfully"})
}

func RejectFriendRequest(c echo.Context) error {
	u, err := usermodel.LoadByRequest(c)
	if err != nil {
		return err
	}
	requester, err := usermodel.Load(c.Param("requester"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	friendReq, err := usermodel.LoadPendingFriendRequest(u.ID, requester.ID)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	friendReq.Rejected = true
	if err := friendReq.Save(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	return c.JSON(200, echo.Map{"message": "friend acccepted successfully"})
}

func FriendRequests(c echo.Context) error {
	u, err := usermodel.LoadByRequest(c)
	if err != nil {
		return err
	}
	return c.JSON(200, echo.Map{
		"pendingList": u.FriendRequestsPendingList(),
	})
}

func FriendList(c echo.Context) error {
	u, err := usermodel.LoadByRequest(c)
	if err != nil {
		return err
	}

	friends := usermodel.Query(bson.M{"id": bson.M{"$in": u.Friends}})

	response := []echo.Map{}
	for _, f := range friends {
		response = append(response, miniResponse(&f))
	}

	return c.JSON(200, echo.Map{
		"friends": response,
	})
}
