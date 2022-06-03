package models

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	InitMysql()
	fmt.Println(db)
	user := User{
		Username: "zhumenglin",
		Password: "456789",
		Nickname: "zml",
	}
	code := AddUser(&user)
	fmt.Println(code)
}

func TestFindUser(t *testing.T) {
	InitMysql()
	user, _ := FindUser(1)
	fmt.Println(user)
}

func TestUpdateUser(t *testing.T) {
	InitMysql()
	user := &User{
		Username: "sunsong",
		Password: "0004456",
		Nickname: "skss",
	}
	UpdateUser(1, user)
}

func TestDeleteUser(t *testing.T) {
	InitMysql()
	DeleteUser(1)
}
