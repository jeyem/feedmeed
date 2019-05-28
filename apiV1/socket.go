package apiV1

import (
	"github.com/Sirupsen/logrus"
	"github.com/jeyem/feedmeed/models/usermodel"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

func SocketConnect(c echo.Context) error {

	var token string

	websocket.Handler(func(ws *websocket.Conn) {
		if err := websocket.Message.Receive(ws, &token); err != nil {
			logrus.Warning("not authorize socket tried")
			ws.Close()
			return
		}
		user, err := usermodel.LoadByToken(token)
		if err != nil {
			logrus.Error(err)
			ws.Close()
			return
		}
		logrus.Info(user.Username, " socket connect successfully !")
		socket := &usermodel.Socket{
			ID:     user.ID,
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
