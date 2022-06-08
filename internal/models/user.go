package models

import (
	"etweb/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"uniqueIndex,type:varchar(20);not null"`
	Password   string `gorm:"type:varchar(500);not null" `
	Nickname   string `gorm:"type:varchar(500)" `
	Profilepic string `gorm:"type:varchar(1024)"`
}

// 新增用户
func AddUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return utils.ERROR
	} else {
		return utils.SUCCSE
	}
}

// 删除用户
func DeleteUser(username string) int {
	var user *User
	var code int
	user, code = FindUser(username)
	if code == utils.ERROR {
		return utils.ERROR
	}
	err := db.Delete(user)
	if err != nil {
		return utils.SUCCSE
	} else {
		return utils.ERROR
	}
}

// 更新用户
func UpdateUser(username string, data *User) int {
	var user *User
	user, _ = FindUser(username)
	if data.Username != user.Username {
		return utils.ERROR
	}
	err = db.Select("password", "nickname", "profilepic").Where("username = ?", username).Updates(&data).Error
	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCSE
}

// 查询用户
func FindUser(username string) (*User, int) {
	var user User
	db.Where("username = ?", username).First(&user)
	// fmt.Println("查询到用户的id", user.ID)
	if user.ID > 0 {
		return &user, utils.SUCCSE
	} else {
		return nil, utils.ERROR
	}
}
