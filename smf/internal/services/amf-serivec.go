package services

import (
	"bytes"
	// "net/http"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/KhanhLinh2810/5G-core/smf/internal/types"
	"github.com/KhanhLinh2810/5G-core/smf/pkg/config"
)

func SendN1N2Mess(sessionInfo *types.CreateSessionRequest) error {
	amfURL := fmt.Sprintf("http://amf:9010/namf-comm/v1/ue-context/%s/n1-n2-messages", sessionInfo.Supi)
	jsonData, err := json.Marshal(sessionInfo)
	if err != nil {
		return fmt.Errorf("failed to marshal sessionInfo: %w", err)
	}

	indexHttp := rand.Intn(config.NUM_CLIENT)
	httpClient := config.ListHttpClient[indexHttp]

	resp, err := httpClient.Post(amfURL, "application/json", bytes.NewBuffer(jsonData))
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
