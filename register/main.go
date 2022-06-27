package main

import (
	"fmt"
	"register/entity"
	"register/service"
	"time"
)

func main() {

	userSvc := service.NewUserSvc()
	if user, err := userSvc.Register(&entity.User{
		Id:        1,
		Username:  "budi123",
		Email:     "budi123@gmail.com",
		Password:  "password123",
		Age:       9,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}); err != nil {
		fmt.Printf("Error when register user: %+v", err)
		return
	} else {
		fmt.Printf("Success register user: %+v", user)
	}

}
