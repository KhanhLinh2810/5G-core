package services

import (
	"fmt"
	// "log"
	"net/http"

	"github.com/google/uuid"
	"github.com/KhanhLinh2810/5G-core/smf/internal/models"
	"github.com/KhanhLinh2810/5G-core/smf/internal/types"
)

func CreateSession(req types.CreateSessionRequest, ResultChan chan any) {
	status,  _, err := ValidateImsi(req.Supi)
	if err != nil {
		ResultChan <- map[string]any{
			"status": status,
			"mess":   err,
		}
		return
	}

	result := map[string]any{
		"status": http.StatusOK,
		"mess":   "Session request accepted",
	}
	ResultChan <- result

	go func() {
		pfcpMsg := &types.PFCPMessage{
			MessageType: 50,
			PDNType:     "IPv4",
			IPAddress:   "10.11.22.123",
			SessionID:   uuid.NewString(),
		}

		if err := SendPFCPJsonUDP(pfcpMsg, "upf:8805"); err != nil {
			fmt.Println("SendPFCPJsonUDP error:", err)
		}

		if err := SendN1N2Mess(&req); err != nil {
			fmt.Println("SendN1N2Mess error:", err)
		}

		if err := models.SaveSession(&req); err != nil {
			fmt.Println("SaveSession error:", err)
		}
	}()
}
