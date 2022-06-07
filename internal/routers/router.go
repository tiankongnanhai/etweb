package routers

import (
	"encoding/gob"
	"etweb/internal/models"
	"etweb/internal/service"
	"etweb/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// cookie存储session
var store = cookie.NewStore([]byte("secret"))

func Initrouter() {
	gob.Register(models.User{}) // 注册User结构体
	gin.SetMode(utils.RunMode)
	r := gin.New()
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("/api/user/")
	{
		v1.POST("/register", service.Register)
		v1.POST("/login", service.Login)
		v1.GET("/get", service.GetUserInfro)
		v1.POST("/edit", service.EditUser)
		v1.POST("/uploadfile", service.UploadFile)
	}
	r.Run(utils.ServerPort)
}
