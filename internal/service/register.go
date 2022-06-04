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
	// 检查username是否已经注册
	isRegister := CheakRegisterExit(username)
	if isRegister {
		c.JSON(
			http.StatusOK, gin.H{
				"msg":     utils.GetErrMsg(utils.USER_REGISTER_FAILD),
				"retcode": "USER_REGISTER_FAILD",
			},
		)
		return
	}
	password := c.PostForm("password")
	nickname := c.PostForm("nickname")
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
