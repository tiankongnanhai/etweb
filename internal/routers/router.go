package routers

import (
	"etweb/internal/service"
	"etweb/utils"

	"github.com/gin-gonic/gin"
)

func Initrouter() {
	gin.SetMode(utils.RunMode)
	r := gin.New()
	v1 := r.Group("/api/user/")
	{
		v1.POST("/register", service.Register)
		v1.GET("/login", service.Login)
	}
	r.Run(utils.ServerPort)
}
