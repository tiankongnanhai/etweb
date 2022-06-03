package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=64" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=4,max=64" label:"密码"`
	Nickname string `gorm:"type:varchar(500);not null" json:"nickname" validate:"required,min=4,max=64" label:"别名"`
}

func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return 1001 //1001
	}
	return 1002
}

func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return 500 // 500
	}
	return 501
}
