package service

import (
	"etweb/internal/models"
	"etweb/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func EditUser(c *gin.Context) {
	// 创建seesion
	session := sessions.Default(c)

	// session过期
	if session.Get("currentUser") == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":     utils.GetErrMsg(utils.USER_EDIT_PROFILE_FAILED),
			"retcode": utils.USER_EDIT_PROFILE_FAILED,
		})
		return
	}

	var userJson models.User
	if err := c.ShouldBindJSON(&userJson); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	oldUser := session.Get("currentUser").(models.User)
	newUser := models.User{
		Username:   oldUser.Username,
		Password:   oldUser.Password,
		Nickname:   userJson.Nickname,
		Profilepic: userJson.Profilepic,
	}
	models.UpdateUser(oldUser.Username, &newUser)
	c.JSON(http.StatusOK, gin.H{
		"data":    struct{}{},
		"msg":     utils.GetErrMsg(utils.SUCCSE),
		"retcode": utils.SUCCSE,
	})
}
