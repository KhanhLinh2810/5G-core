package services

import (
	"fmt"
	"net/http"
	"io/ioutil"

	"github.com/KhanhLinh2810/5G-core/smf/pkg/config"

)

func ValidateImsi(supi string) (int, []byte, error) {
	udmURL := fmt.Sprintf("http://udm:8000/nudm-sdm/v2/%s/sm-data", supi)

	resp, err := config.HttpClient.Get(udmURL)
	if err != nil {
		return http.StatusBadGateway, nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode == http.StatusNotFound {
		return resp.StatusCode, body, fmt.Errorf("smf not found %s", supi)
	} 
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, body, fmt.Errorf("server returned status %d: %s", resp.StatusCode, string(body))
	}

	return resp.StatusCode, body, nil
}