package routers

import (
	"etweb/utils"

	"github.com/gin-gonic/gin"
)

func Initrouter() {
	gin.SetMode(utils.RunMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(utils.ServerPort)
}
