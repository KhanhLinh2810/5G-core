package main

import (
	"github.com/KhanhLinh2810/5G-core/udm/pkg/routes"
	"github.com/KhanhLinh2810/5G-core/udm/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectRedis()
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	routes.SMFRoutes(router)
	router.Run(":8000") 
}
