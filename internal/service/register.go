package service

import (
	"etweb/internal/models"
	"etweb/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	// 从body中获取username和password
	username := c.PostForm("username")
	password := c.PostForm("password")
	// username为空，返回
	if username == "" {
		c.JSON(
			http.StatusOK, gin.H{
				"msg":     utils.GetErrMsg(utils.USERNAME_EMPTY),
				"retcode": utils.USERNAME_EMPTY,
			},
		)
		return
	}
	// password为空，返回
	if password == "" {
		c.JSON(
			http.StatusOK, gin.H{
				"msg":     utils.GetErrMsg(utils.PASSWORD_EMPTY),
				"retcode": utils.PASSWORD_EMPTY,
			},
		)
		return
	}
	// 检查username是否已经注册
	isRegister := CheakRegisterExit(username)
	if isRegister {
		c.JSON(
			http.StatusOK, gin.H{
				"msg":     utils.GetErrMsg(utils.USER_REGISTER_FAILD),
				"retcode": utils.USER_REGISTER_FAILD,
			},
		)
		return
	}
	// 获取nickname，可以为空
	nickname := c.PostForm("nickname")
	// 写入数据库
	models.AddUser(&models.User{Username: username, Password: password, Nickname: nickname})
	c.JSON(http.StatusOK, gin.H{
		"data":    struct{}{},
		"msg":     utils.GetErrMsg(utils.SUCCSE),
		"retcode": utils.SUCCSE,
	})

}

func CheakRegisterExit(username string) bool {
	user, _ := models.FindUser(username)
	if user != nil {
		return true
	} else {
		return false
	}
}
