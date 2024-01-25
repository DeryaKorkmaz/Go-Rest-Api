package main

import (
	"myproject/config"
	"myproject/routes"

	"github.com/gin-gonic/gin"
	// "github.com/rs/cors/wrapper/gin"
)

func main() {

	router := gin.New()
	config.Connect()

	routes.UserRoute(router)

	router.Run(":8080")

}
