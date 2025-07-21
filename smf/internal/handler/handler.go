// package handler

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"

// 	"github.com/gin-gonic/gin"

// 	"github.com/KhanhLinh2810/5G-core/amf/internal/service/ue_service" 
// 	"github.com/KhanhLinh2810/5G-core/amf/internal/service/smf_service" 
// )

// func CreateSession(c *gin.Context) {
// 	var req CreateSessionRequest

// 	// Bind JSON request body to CreateSessionRequest struct
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, Response{
// 			Status:  "error",
// 			Message: fmt.Sprintf("Invalid request: %v", err),
// 		})
// 		return
// 	}

// 	// Validate required fields (additional validation can be added here)
// 	if req.supi == "" || req.pduSessionId <= 0 || req.dnn == "" {
// 		c.JSON(http.StatusBadRequest, Response{
// 			Status:  "error",
// 			Message: "Missing or invalid required fields",
// 		})
// 		return
// 	}

// 	// Simulate SMF processing (e.g., validate IMSI with UDM, send PFCP to UPF, etc.)
// 	// In a real system, you would add logic to:
// 	// - Call UDM for IMSI validation
// 	// - Send PFCP Session Establishment to UPF
// 	// - Store session in database
// 	// - Return N1N2 Message Transfer to AMF

// 	// For this example, return a success response
// 	c.JSON(http.StatusOK, Response{
// 		Status:  "success",
// 		Message: fmt.Sprintf("PDU Session created for IMSI: %s, PDU Session ID: %d", req.Supi, req.PduSessionId),
// 	})
// }