package service

import (
	"fmt"
	"register/entity"
)

type UserServiceIface interface {
	Register(user *entity.User) (users *entity.User, err error)
}

type UserSvc struct {
	user entity.User
}

func NewUserSvc() UserServiceIface {
	return &UserSvc{}
}

func (u *UserSvc) Register(user *entity.User) (users *entity.User, err error) {
	fmt.Println(user)
	return user, nil
}
