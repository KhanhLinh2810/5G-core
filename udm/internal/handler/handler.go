// package handler

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"

// 	"github.com/gin-gonic/gin"

// 	"github.com/KhanhLinh2810/5G-core/amf/internal/service/ue_service" 
// 	"github.com/KhanhLinh2810/5G-core/amf/internal/service/smf_service" 
// )

// func GetSDMData(c *gin.Context) {
// 	imsi := c.Param("imsi")

// 	// Validate IMSI
// 	if imsi == "" {
// 		c.JSON(http.StatusBadRequest, Response{
// 			Status:  "error",
// 			Message: "IMSI is required",
// 		})
// 		return
// 	}

// 	// Query Redis for SM data
// 	ctx := context.Background()
// 	key := fmt.Sprintf("imsi:%s", imsi)
// 	data, err := rdb.HGetAll(ctx, key).Result()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, Response{
// 			Status:  "error",
// 			Message: fmt.Sprintf("Failed to query Redis: %v", err),
// 		})
// 		return
// 	}

// 	// Check if data exists
// 	if len(data) == 0 {
// 		c.JSON(http.StatusNotFound, Response{
// 			Status:  "error",
// 			Message: fmt.Sprintf("No SM data found for IMSI: %s", imsi),
// 		})
// 		return
// 	}

// 	// Convert pduSessionId from string to int
// 	var pduSessionId int
// 	if pduSessionIdStr, ok := data["pduSessionId"]; ok {
// 		_, err := fmt.Sscanf(pduSessionIdStr, "%d", &pduSessionId)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, Response{
// 				Status:  "error",
// 				Message: fmt.Sprintf("Invalid pduSessionId format: %v", err),
// 			})
// 			return
// 		}
// 	}

// 	// Construct SMData response
// 	smData := SMData{
// 		Imsi:         imsi,
// 		Dnn:          data["dnn"],
// 		PduSessionId: pduSessionId,
// 		SNssai: struct {
// 			Sst string `json:"sst"`
// 			Sd  string `json:"sd"`
// 		}{
// 			Sst: data["sst"],
// 			Sd:  data["sd"],
// 		},
// 	}

// 	// Return success response with SM data
// 	c.JSON(http.StatusOK, Response{
// 		Status:  "success",
// 		Message: "SM data retrieved successfully",
// 		Data:    smData,
// 	})
// }