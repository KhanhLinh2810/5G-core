package services

import (
	"bytes"
	"net/http"
	"fmt"
	"encoding/json"

	"github.com/KhanhLinh2810/5G-core/smf/internal/types"
)

func SendN1N2Mess(sessionInfo *types.CreateSessionRequest) error {
	amfURL := fmt.Sprintf("http://amf:9010/namf-comm/v1/ue-context/%s/n1-n2-messages", sessionInfo.Supi )
	jsonData, err := json.Marshal(sessionInfo)
	if err != nil {
		return fmt.Errorf("failed to marshal sessionInfo: %w", err)
	}

	resp, err := http.Post(amfURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("HTTP POST failed: %w", err)
	}
	defer resp.Body.Close()

	// Kiểm tra status code
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil // Thành công
	}

	return fmt.Errorf("AMF returned status %d", resp.StatusCode)
}