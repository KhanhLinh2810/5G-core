package main

import (
	"github.com/KhanhLinh2810/5G-core/amf/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	routes.UESessionRoutes(router)
	router.Run(":9010") 
}
