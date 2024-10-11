package app

import (
	"errors"

	"github.com/aliasthewho/online-store/pkg/users"
)

type UserService struct {
	users []users.User
}

func NewUserService() *UserService {
	return &UserService{
		users: []users.User{
			{ID: 1, Name: "Alice", Email: "a@a"},
			{ID: 2, Name: "Bob", Email: "b@b"},
		},
	}
}

func (us *UserService) ListUsers() []users.User {
	return us.users
}

func (us *UserService) GetUserByID(id int) *users.User {
	for _, u := range us.users {
		if u.ID == id {
			return &u
		}
	}
	return nil
}

func (us *UserService) CreateUser(user users.User) {
	us.users = append(us.users, user)
}

func (us *UserService) UpdateUser(id int, updateUser users.User) error {
	for i, u := range us.users {
		if u.ID == id {
			us.users[i] = updateUser
			return nil
		}
	}
	return errors.New("user not found")
}

func (us *UserService) DeleteUser(id int) error {
	for i, u := range us.users {
		if u.ID == id {
			us.users = append(us.users[:i], us.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
