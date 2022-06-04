package models

import (
	"etweb/utils"
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
	user, code := FindUser("pengxiangrong")
	fmt.Println(code)
	if code == utils.SUCCSE {
		fmt.Println(user, user.Username, user.Password)
	} else {
		fmt.Println("not record")
	}
}

func TestUpdateUser(t *testing.T) {
	InitMysql()
	user := &User{
		Username: "sunsong",
		Password: "0004456",
		Nickname: "skss",
	}
	UpdateUser("sunsong", user)
}

func TestDeleteUser(t *testing.T) {
	InitMysql()
	DeleteUser("sunsong")
}
