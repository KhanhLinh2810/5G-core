package controllers

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/KhanhLinh2810/5G-core/amf/internal/services"
)

func UECreateSession(c *gin.Context) {
	csrJSON := services.MockDataForUERequestHandler()
	
	// // send to SMF
	// resp, err := services.CreateSession(csrJSON)
	// if err != nil {
	// 	c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to send to SMF"})
	// 	return
	// }
	// defer resp.Body.Close()

	// // Log response from SMF
	// log.Println("AMF received response from SMF, status:", resp.Status)

	// // Respond to client
	// c.JSON(http.StatusOK, gin.H{"status": "Session request processed", "smfStatus": resp.Status})
		
	
	log.Println("AMF received response from SMF, status:", csrJSON)

	c.JSON(http.StatusOK, gin.H{"status": "Session request processed", "smfStatus": 200})
}