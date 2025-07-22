package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	
	"github.com/KhanhLinh2810/5G-core/amf/internal/types"

)

func MockDataForUERequestHandler() []byte {
	csr := types.TypeCreateSessionRequest{
		Supi:        "imsi-20893000000085",
		Gpsi:        "msisdn-84900000001",
		PduSessionID: 1,
		Dnn:         "iot",
		ServingNHd:  "2ab2b5a9-68e8-4ee6-b939-024c109b520c",
		AnType:      "3GPP_ACCESS",
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
	smfURL := "http://localhost:40/nsmf-pdusession/v1/sm-contexts"
	resp, err := http.Post(smfURL, "application/json", bytes.NewBuffer(csrJSON))
	if err != nil {
		return nil, err
	}
	return resp, nil
}