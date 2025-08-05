package main

import (
	"github.com/KhanhLinh2810/5G-core/smf/pkg/config"
	"github.com/KhanhLinh2810/5G-core/smf/pkg/routes"

	// "github.com/KhanhLinh2810/5G-core/smf/internal/models"
	"github.com/KhanhLinh2810/5G-core/smf/internal/workers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectRedis()
	config.InitHTTPClient()
	gin.SetMode(gin.ReleaseMode)

	go workers.StartWorkerPool(100)

	router := gin.New()
	routes.AMFRoutes(router)
	router.Run(":40")
}
