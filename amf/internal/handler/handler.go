package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KhanhLinh2810/5G-core/amf/internal/service/ue_service" 
	"github.com/KhanhLinh2810/5G-core/amf/internal/service/smf_service" 
)

func UERequestHandler(c *gin.Context) {
	csrJSON = MockDataForUERequestHandler()
	
	// Gửi tới SMF
	resp, err := smf_service.CreateSession(csrJSON)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to send to SMF"})
		return
	}
	defer resp.Body.Close()

	// Log response from SMF
	c.Logger.Println("AMF received response from SMF, status:", resp.Status)

	// Respond to client
	c.JSON(http.StatusOK, gin.H{"status": "Session request processed", "smfStatus": resp.Status})
}