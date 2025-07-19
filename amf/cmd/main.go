package main

import (
	"os"
	"fmt"
	"log"
	"net/http"

	"github.com/KhanhLinh2810/5G-core/amf/internal/type" 
	"github.com/KhanhLinh2810/5G-core/amf/internal/handler" 

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	// "github.com/caarlos0/env/v6" 
)

func main() {
	// Load env file
	fmt.Println("Loading env...")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
	cfg := Config{}

	baseUrl := os.Getenv("AMF_BASE_URL")
	port := os.Getenv("AMF_PORT")

	addr := fmt.Sprintf("%s:%s", baseUrl, port)

	// Initialize router 
	router := gin.Default()

	// API
	router.GET("/health-amf", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is starting",
		})
	})

	router.GET("/ue-session", handler.UERequestHandler)

	// server running
	router.Run(addr)

}