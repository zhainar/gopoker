package model

import (
	"github.com/gorilla/websocket"
)

const (
	connected = "connected"
	offline   = "offline"
)

type User struct {
	conn   *websocket.Conn
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func NewUser(conn *websocket.Conn, data map[string]string) *User {
	return &User{
		conn:   conn,
		ID:     data["id"],
		Name:   data["name"],
		Status: connected,
	}
}

func (u *User) Notify(data interface{}) error {
	return u.conn.WriteJSON(data)
}
