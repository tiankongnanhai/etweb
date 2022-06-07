package service

import (
	"etweb/internal/models"
	"etweb/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetUserInfro(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("currentUser") != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":     session.Get("currentUser").(models.User),
			"retcode": utils.ERROR,
		})
	} else {
		c.JSON(http.StatusBadRequest, "登陆过期或未登录")
	}
}
