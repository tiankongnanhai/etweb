package service

import (
	"etweb/internal/models"
	"etweb/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 设置seesion
	session := sessions.Default(c)
	if session.Get("hello") != "world" {
		session.Set("hello", "world")
	}
	session.Save()
	// 检查username是否已经注册
	// 如果没注册，不允许登陆，返回
	isRegister := CheakRegisterExit(username)
	if !isRegister {
		c.JSON(
			http.StatusOK, gin.H{
				"msg":     utils.GetErrMsg(utils.USER_NOT_REGISTER),
				"retcode": utils.USER_NOT_REGISTER,
			},
		)
		return
	}
	// 检查密码是否正确
	// 正确：登陆
	// 错误：返回
	user, _ := models.FindUser(username)
	if password != user.Password {
		c.JSON(
			http.StatusOK, gin.H{
				"msg":     utils.GetErrMsg(utils.USER_LOGIN_FAILD),
				"retcode": utils.USER_LOGIN_FAILD,
			},
		)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"sessionId": session.ID(),
			"msg":       utils.GetErrMsg(utils.SUCCSE),
			"retcode":   utils.SUCCSE,
		})
	}
}
