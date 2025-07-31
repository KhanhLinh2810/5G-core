package main

import (
	"github.com/KhanhLinh2810/5G-core/smf/pkg/routes"
	"github.com/KhanhLinh2810/5G-core/smf/pkg/config"
	"github.com/KhanhLinh2810/5G-core/smf/internal/models"
	// "github.com/KhanhLinh2810/5G-core/smf/internal/worker"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectRedis() 
	models.SeedFakeSessions()

	// go worker.StartWorkerPool(5)

	router := gin.Default()
	routes.AMFRoutes(router)
	router.Run(":40") 
}
