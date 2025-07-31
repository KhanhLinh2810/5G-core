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

func CreateSessionSaveInMap(req types.CreateSessionRequest) {
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

		if err := services.SendPFCPJsonUDP(pfcpMsg, "127.0.0.1:8805"); err != nil {
			fmt.Println("SendPFCPJsonUDP error:", err)
		}

		if err := services.SendN1N2Mess(&req); err != nil {
			fmt.Println("SendN1N2Mess error:", err)
		}

		session := models.Session{
			Supi:         req.Supi,
			Gpsi:         req.Gpsi,
			PduSessionID: req.PduSessionID,
			Dnn:          req.Dnn,
			Action:       req.Action,
		}

		models.GlobalSessionStore.SaveSessionInMap(session)
	}()
}

func ReleaseSession(req types.ReleaseSessionRequest) {
		_, found := store.GlobalSessionStore.Get(req.Supi)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}

	model.GlobalSessionStore.Delete(req.Supi)
	log.Printf("[Release] Deleted session for SUPI: %s", req.Supi)

	c.JSON(http.StatusOK, gin.H{
		"message": "session released",
		"supi":    req.Supi,
	})
}

func UpdateSession(req types.Session) {
		session, found := store.GlobalSessionStore.Get(req.Supi)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}

	isValidToUpdate := CheckValidActionToUpdate(req.Action, session.Action)
	if !isValidToUpdate {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "false",
			"supi":    req.Supi,
			"old":     session.Action,
			"new":     req.Action,
		})
		return
	}

	models.GlobalSessionStore.SaveSessionInMap(req) 

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"supi":    req.Supi,
		"old":     session.Action,
		"new":     req.Action,
	})
}


// other function
func CheckValidActionToUpdate(actionNew, actionOld string) bool {
	allowed := map[string][]string{
		"ACTIVE":   {"DEACTIVE", "HANDOVER", "CALL"},
		"DEACTIVE": {"ACTIVE", "HANDOVER"},
		"HANDOVER": {"DEACTIVE", "HANDOVER", "ACTIVE"},
		"CALL":     {"DEACTIVE", "HANDOVER", "ACTIVE"},
	}

	validNextActions, ok := allowed[actionOld]
	if !ok {
		return false 
	}

	for _, a := range validNextActions {
		if a == actionNew {
			return true
		}
	}
	return false
}

