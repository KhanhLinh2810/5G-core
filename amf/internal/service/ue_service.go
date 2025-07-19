package ue_service

import (
	"encoding/json"

	"github.com/KhanhLinh2810/5G-core/amf/internal/type"
)

func MockDataForUERequestHandler() []byte {
	csr := type.TypeCreateSessionRequest{
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