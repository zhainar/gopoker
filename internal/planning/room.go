package planning

import (
	"github.com/zhainar/gopoker/internal/model"
)

type Room struct {
	users []*model.User
}

func NewRoom() *Room {
	return &Room{}
}

func (r *Room) AddUser(u *model.User) {
	if r.UserConnected(u.ID) {
		r.UpdateUser(u)
	} else {
		r.users = append(r.users, u)
	}
}

func (r *Room) DetachUser(user *model.User) {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = r.users[len(r.users)-1]
			r.users = r.users[:len(r.users)-1]
		}
	}
}

func (r *Room) Notify(data interface{}) error {
	for _, u := range r.users {
		err := u.Notify(data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Room) UpdateUsers() {
	r.Notify(map[string]interface{}{
		"users": r.users,
	})
}

func (r *Room) UserConnected(id string) bool {
	for _, u := range r.users {
		if u.ID == id {
			return true
		}
	}

	return false
}

func (r *Room) UpdateUser(user *model.User) {
	for i, u := range r.users {
		if u.ID == u.ID {
			r.users[i] = user
		}
	}
}
