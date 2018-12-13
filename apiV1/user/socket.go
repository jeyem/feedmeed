package user

import (
	"github.com/Sirupsen/logrus"
	"github.com/jeyem/feedmeed/models/usermodel"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
	"gopkg.in/mgo.v2/bson"
)

func SocketConnect(c echo.Context) error {

	var token string

	websocket.Handler(func(ws *websocket.Conn) {
		if err := websocket.Message.Receive(ws, &token); err != nil {
			logrus.Warning("not authorize socket tried")
			ws.Close()
			return
		}
		id, err := usermodel.GetUsernameFromToken(token)
		if err != nil {
			logrus.Warning("failed load token from cache")
			ws.Close()
			return
		}

		socket := &usermodel.Socket{
			ID:     bson.ObjectIdHex(id),
			Token:  token,
			Caster: make(chan *usermodel.BroadCaster, 8),
		}

		usermodel.Connections.New(socket)

		for c := range socket.Caster {
			websocket.Message.Send(ws, c.Message())
		}

		ws.Close()
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
