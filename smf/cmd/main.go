package main

import (
	"github.com/KhanhLinh2810/5G-core/smf/pkg/config"
	"github.com/KhanhLinh2810/5G-core/smf/pkg/routes"

	// "github.com/KhanhLinh2810/5G-core/smf/internal/models"
	"net/http"

	"github.com/KhanhLinh2810/5G-core/smf/internal/workers"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	config.ConnectRedis()
	config.InitHTTPClient()
	gin.SetMode(gin.ReleaseMode)

	go workers.StartWorkerPool(100)

	router := gin.New()
	routes.AMFRoutes(router)

	h2s := &http2.Server{}
	handler := h2c.NewHandler(router, h2s)

	addr := ":40"
	// Use custom HTTP server
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
