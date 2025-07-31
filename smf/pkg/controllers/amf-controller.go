package controllers

import (
	"net/http"
	"fmt"
	// "sync"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/KhanhLinh2810/5G-core/smf/internal/types"
	"github.com/KhanhLinh2810/5G-core/smf/internal/models"
	"github.com/KhanhLinh2810/5G-core/smf/internal/services"
)

func AMFCreateSession(c *gin.Context) {
	var req types.CreateSessionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	// Validate SUPI với UDM
	body, err := services.ValidateImsi(req.Supi)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": fmt.Sprintf("Failed to validate SUPI with UDM: %v", err),
		})
		return
	}
	fmt.Printf("response of udm: %s\n", body)

	// Trả kết quả luôn cho client
	c.JSON(http.StatusOK, gin.H{
		"status": "Session request accepted",
	})

	go func() {
		pfcpMsg := &types.PFCPMessage{
			MessageType: 50,
			PDNType:     "IPv4",
			IPAddress:   "10.11.22.123",
			SessionID:   uuid.NewString(),
		}

		if err := services.SendPFCPJsonUDP(pfcpMsg, "upf:8805"); err != nil {
			fmt.Println("SendPFCPJsonUDP error:", err)
		}

		if err := services.SendN1N2Mess(&req); err != nil {
			fmt.Println("SendN1N2Mess error:", err)
		}

		// Redis
		err := models.SaveSession(&req)
		if err != nil {
		    fmt.Println("SaveSession error:", err)
		}
	}()
}
