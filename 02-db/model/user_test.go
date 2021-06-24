package model

import (
	"fmt"
	"testing"
)

func TestUser_Save(t *testing.T) {
	user := User{Username: "mary", Email: "aaa@g.com"}
	err := user.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestUser_GetById(t *testing.T) {
	user := User{Id: 1}
	err := user.GetById()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
}
