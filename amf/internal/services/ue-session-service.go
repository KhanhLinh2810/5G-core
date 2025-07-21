package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	
	"github.com/KhanhLinh2810/5G-core/amf/internal/types"

)

func MockDataForUERequestHandler() []byte {
	csr := types.TypeCreateSessionRequest{
		Supi:        "imsi-452040000000001",
		Gpsi:        "msisdn-84900000001",
		PduSessionID: 1,
		Dnn:         "v-internet",
		ServingNHd:  "2ab2b5a9-68e8-4ee6-b939-024c109b520c",
		AnType:      "3GPP_ACCESS",
	}
	csr.SNssai.Sst = 1
	csr.SNssai.Sd = "000001"

	csrJSON, err := json.Marshal(csr)
	if err != nil {
		return nil
	}
	return csrJSON
}

func CreateSession(csrJSON []byte) (*http.Response, error) {
	smfURL := "http://localhost:40/nsmf-pdusession/v1/sm-contexts"
	resp, err := http.Post(smfURL, "application/json", bytes.NewBuffer(csrJSON))
	if err != nil {
		return nil, err
	}
	return resp, nil
}