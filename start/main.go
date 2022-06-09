package main

import (
	"etweb/internal/models"
	"etweb/internal/routers"
)

func main() {
	// mysql setting
	models.InitMysql()
	// router setting
	routers.Initrouter()
}
