package main

import (
	"net/http"

	"github.com/KhanhLinh2810/5G-core/udm/pkg/config"
	"github.com/KhanhLinh2810/5G-core/udm/pkg/routes"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	config.ConnectRedis()
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	routes.SMFRoutes(router)

	h2s := &http2.Server{}
	handler := h2c.NewHandler(router, h2s)

	addr := ":8000"
	// Use custom HTTP server
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
