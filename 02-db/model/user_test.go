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
