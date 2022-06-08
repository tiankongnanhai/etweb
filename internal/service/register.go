package service

import (
	"etweb/internal/models"
	"etweb/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var userJson models.User
	if err := c.ShouldBindJSON(&userJson); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// username为空，返回
	if userJson.Username == "" {
		c.JSON(
			http.StatusOK, gin.H{
				"msg":     utils.GetErrMsg(utils.USERNAME_EMPTY),
				"retcode": utils.USERNAME_EMPTY,
			},
		)
		return
	}
	// password为空，返回
	if userJson.Password == "" {
		c.JSON(
			http.StatusOK, gin.H{
				"msg":     utils.GetErrMsg(utils.PASSWORD_EMPTY),
				"retcode": utils.PASSWORD_EMPTY,
			},
		)
		return
	}
	// 检查username是否已经注册
	isRegister := CheakRegisterExit(userJson.Username)
	if isRegister {
		c.JSON(
			http.StatusOK, gin.H{
				"msg":     utils.GetErrMsg(utils.USER_REGISTER_FAILD),
				"retcode": utils.USER_REGISTER_FAILD,
			},
		)
		return
	}
	// 写入数据库
	models.AddUser(&userJson)
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
