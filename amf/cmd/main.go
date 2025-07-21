package main

import (
	"github.com/KhanhLinh2810/5G-core/amf/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.UESessionRoutes(router)
	router.Run("localhost:9010") 
}
