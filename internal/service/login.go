package service

import (
	"etweb/internal/models"
	"etweb/utils"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userJson models.User
	if err := c.ShouldBindJSON(&userJson); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查username是否已经注册
	// 如果没注册，不允许登陆，返回
	isRegister := CheakRegisterExit(userJson.Username)
	if !isRegister {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"msg":     utils.GetErrMsg(utils.USER_NOT_REGISTER),
				"retcode": utils.USER_NOT_REGISTER,
			},
		)
		return
	}

	// 检查密码是否正确
	// 正确：登陆
	// 错误：返回
	user, _ := models.FindUser(userJson.Username)
	if userJson.Username != user.Password {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"msg":     utils.GetErrMsg(utils.USER_LOGIN_FAILD),
				"retcode": utils.USER_LOGIN_FAILD,
			},
		)
		return
	}

	// 创建seesion
	session := sessions.Default(c)

	// session中储存当前用户
	v := session.Get("currentUser")
	if v != user {
		session.Set("currentUser", user)
	}

	// 设置session的参数
	options := sessions.Options{}
	options.MaxAge = 100 // 有效期10秒
	session.Options(options)
	session.Save()

	fmt.Println(session.ID())

	// 正确，返回
	c.JSON(http.StatusOK, gin.H{
		"sessionId": session.ID(),
		"msg":       utils.GetErrMsg(utils.SUCCSE),
		"retcode":   utils.SUCCSE,
	})

}
