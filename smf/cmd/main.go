package main

import (
	"github.com/KhanhLinh2810/5G-core/smf/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.AMFRoutes(router)
	router.Run("localhost:40") 
}
