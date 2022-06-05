package routers

import (
	"etweb/internal/service"
	"etweb/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// redis存储session
var store, _ = redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))

func Initrouter() {
	gin.SetMode(utils.RunMode)
	r := gin.New()
	r.Use(sessions.Sessions("sessionId", store))
	v1 := r.Group("/api/user/")
	{
		v1.POST("/register", service.Register)
		v1.POST("/login", service.Login)
	}
	r.Run(utils.ServerPort)
}
