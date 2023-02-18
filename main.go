package main

import (
	"user-center/config"
	"user-center/routes"
)

func main() {
	err := routes.R.Run(config.Config.Gin.Address)
	if err != nil {
		return
	}
}
