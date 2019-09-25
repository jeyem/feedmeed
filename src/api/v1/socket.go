package v1

import (
	"github.com/jeyem/feedmeed/src/models/user"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

func socket(c echo.Context) error {

	var token string

	websocket.Handler(func(ws *websocket.Conn) {
		if err := websocket.Message.Receive(ws, &token); err != nil {
			logrus.Warning("not authorize socket tried")
			ws.Close()
			return
		}
		u, err := user.LoadByToken(token)
		if err != nil {
			logrus.Error(err)
			ws.Close()
			return
		}
		logrus.Info(u.Username, " socket connect successfully !")
		socket := &user.Socket{
			ID:     u.ID,
			Token:  token,
			Caster: make(chan *user.BroadCaster, 8),
		}

		user.Connections.New(socket)

		for c := range socket.Caster {
			websocket.Message.Send(ws, c.Message())
		}

		ws.Close()
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
