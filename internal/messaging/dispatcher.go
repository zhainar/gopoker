package messaging

import (
	"github.com/gorilla/websocket"
	"github.com/zhainar/gopoker/internal/model"
	"github.com/zhainar/gopoker/internal/planning"
)

type Dispatcher struct {
	room *planning.Room
}

func (d *Dispatcher) Handle(conn *websocket.Conn, msg *Message) error {
	switch msg.Action {
	case "connect":
		u := model.NewUser(conn, msg.Data)
		d.room.AddUser(u)
		u.Notify(map[string]string{"result": "connected"})
		d.room.UpdateUsers()
		break
	case "disconnect":
		u := model.NewUser(conn, msg.Data)
		d.room.DetachUser(u)
		u.Notify(map[string]string{"result": "disconnected"})
		d.room.UpdateUsers()
		conn.Close()
		break
	}

	return nil
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		room: planning.NewRoom(),
	}
}
