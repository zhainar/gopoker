package messaging

import (
	"encoding/json"
)

type Message struct {
	Action string            `json:"action"`
	Data   map[string]string `json:"data"`
}

func NewMessage(msg []byte) (*Message, error) {
	var message = &Message{}

	json.Unmarshal(msg, message)

	return message, nil
}
