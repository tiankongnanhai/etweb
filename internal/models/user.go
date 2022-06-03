package models

import (
	"etweb/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=64" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=4,max=64" label:"密码"`
	Nickname string `gorm:"type:varchar(500);not null" json:"nickname" validate:"required,min=4,max=64" label:"别名"`
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
func DeleteUser(id int) int {
	var user *User
	var code int
	user, code = FindUser(id)
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
func UpdateUser(id int, data *User) int {
	var user *User
	user, _ = FindUser(id)
	if data.Username != user.Username {
		return utils.ERROR
	}
	err = db.Select("password", "nickname").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCSE
}

// 查询用户
func FindUser(id int) (*User, int) {
	var user User
	err := db.Limit(1).Where("ID=?", id).Find(&user).Error
	if err != nil {
		return nil, utils.ERROR
	} else {
		return &user, utils.SUCCSE
	}
}
