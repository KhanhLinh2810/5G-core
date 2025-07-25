package controllers

import (
	"net/http"
	"log"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/KhanhLinh2810/5G-core/amf/internal/services"
	"github.com/KhanhLinh2810/5G-core/amf/internal/types"
)

func UECreateSession(c *gin.Context) {
	// Tạo mock request
	csrJSON := services.MockDataForUERequestHandler()

	// Gửi request tới SMF
	resp, err := services.CreateSession(csrJSON)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": fmt.Sprintf("Failed to send to SMF: %v", err),
		})
		return
	}
	if resp == nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "SMF response is nil"})
		return
	}
	defer resp.Body.Close()

	// Log status
	log.Println("AMF received response from SMF, status:", resp.Status)

	// Trả response về client
	c.JSON(http.StatusOK, gin.H{
		"status":    "Session request processed",
		"smfStatus": resp.Status,
	})
}

func N1N2MessageTransfer(c *gin.Context) {
	var req types.N1N2MessageTransfer

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mess": "success recieve n1n2 message",
	})
}
