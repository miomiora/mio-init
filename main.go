package main

import (
	"user-center/config"
	"user-center/routes"
)

func main() {
	err := routes.R.Run(":" + config.Config.Gin.Port)
	if err != nil {
		return
	}
}
