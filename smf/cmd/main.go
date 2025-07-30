package main

import (
	"github.com/KhanhLinh2810/5G-core/smf/pkg/routes"
	"github.com/KhanhLinh2810/5G-core/smf/pkg/config"
	"github.com/KhanhLinh2810/5G-core/smf/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectRedis() 

	models.SeedFakeSessions()

	router := gin.Default()
	routes.AMFRoutes(router)
	router.Run(":40") 
}
