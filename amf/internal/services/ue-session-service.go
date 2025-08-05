package services

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/KhanhLinh2810/5G-core/amf/internal/types"
	"github.com/KhanhLinh2810/5G-core/smf/pkg/config"
)

func MockDataForUERequestHandler() []byte {
	csr := types.TypeCreateSessionRequest{
		// Supi:        "imsi-20893000000085",
		Supi:         "imsi-20893000000026",
		Gpsi:         "msisdn-84900000001",
		PduSessionID: 1,
		Dnn:          "v-internet",
		ServingNHd:   "2ab2b5a9-68e8-4ee6-b939-024c109b520c",
		AnType:       "3GPP_ACCESS",
	}
	csr.SNssai.Sst = 2
	csr.SNssai.Sd = "112233"

	csrJSON, err := json.Marshal(csr)
	if err != nil {
		return nil
	}
	return csrJSON
}

func CreateSession(csrJSON []byte) (*http.Response, error) {
	smfURL := "http://smf:40/nsmf-pdusession/v1/sm-contexts"
	resp, err := config.HttpClient.Post(smfURL, "application/json", bytes.NewBuffer(csrJSON))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// // update
// func MockDataForUpdateUERequestHandler(action string) []byte {
// 	csr := types.UpdateSession{
// 		// Supi:        "imsi-20893000000085",
// 		Supi:        "imsi-20893000000054",
// 		Gpsi:        "msisdn-84900000001",
// 		PduSessionID: 1,
// 		Dnn:         "v-internet",
// 		action:  action,
// 	}

// 	csrJSON, err := json.Marshal(csr)
// 	if err != nil {
// 		return nil
// 	}
// 	return csrJSON
// }

// func UpdateSession(action string) (*http.Response, error) {
// 	csrJSON := services.MockDataForUpdateUERequestHandler(action)

// 	smfURL := "http://smf:40/nsmf-pdusession/v1/update-sm-contexts"
// 	resp, err := http.Post(smfURL, "application/json", bytes.NewBuffer(csrJSON))
// 	if err != nil {
// 		return nil, err
// 	}
// 	bodyBytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Printf("HTTP request failed: %v", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusForbidden {
// 		var data map[string]interface{}
// 		if err := json.Unmarshal(bodyBytes, &data); err != nil {
// 			log.Printf("Failed to parse JSON response: %v", err)
// 			return
// 		}

// 		fmt.Printf("Message: %v\n", data["message"])
// 		fmt.Printf("SUPI: %v\n", data["supi"])
// 		fmt.Printf("Old Action: %v\n", data["old"])
// 		fmt.Printf("New Action: %v\n", data["new"])
// 	} else {
// 		log.Printf("Unexpected status code: %d\nBody: %s", resp.StatusCode, string(bodyBytes))
// 	}
// }

// // release
// func MockDataForReleaseUERequestHandler(action string) []byte {
// 	csr := types.Session{
// 		// Supi:        "imsi-20893000000085",
// 		Supi:        "imsi-20893000000054",
// 		Gpsi:        "msisdn-84900000001",
// 		PduSessionID: 1,
// 		Dnn:         "v-internet",
// 	}

// 	csrJSON, err := json.Marshal(csr)
// 	if err != nil {
// 		return nil
// 	}
// 	return csrJSON
// }

// func ReleaseSession(csrJSON []byte) (*http.Response, error) {
// 	smfURL := "http://smf:40/nsmf-pdusession/v1/release-sm-contexts"
// 	resp, err := http.Post(smfURL, "application/json", bytes.NewBuffer(csrJSON))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp, nil
// }
