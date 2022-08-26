package main

import (
	"SP/dataBase"
	"SP/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	dataBase.DB()
	routes.AuthRoutes(router)
	routes.UserRoute(router)
	router.Run(":6000")
}
