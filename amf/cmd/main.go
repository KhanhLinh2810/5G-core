package main

import (
	"github.com/KhanhLinh2810/5G-core/amf/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/KhanhLinh2810/5G-core/amf/pkg/config"

)

func main() {
	config.InitHTTPClient()
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	routes.UESessionRoutes(router)
	router.Run(":9010")
}
