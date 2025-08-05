package main

import (
	"github.com/KhanhLinh2810/5G-core/amf/pkg/config"
	"github.com/KhanhLinh2810/5G-core/amf/pkg/routes"
	"github.com/gin-gonic/gin"

	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	config.InitHTTPClient()
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	routes.UESessionRoutes(router)

	h2s := &http2.Server{}
	handler := h2c.NewHandler(router, h2s)

	addr := ":9010"
	// Use custom HTTP server
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
