package services

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func ValidateImsi(supi string) ([]byte, error) {
	// Tạo URL hợp lệ
	udmURL := fmt.Sprintf("http://localhost:8000/nudm-sdm/v2/%s/sm-data", supi)

	// Gửi request GET
	resp, err := http.Get(udmURL)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned status %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}