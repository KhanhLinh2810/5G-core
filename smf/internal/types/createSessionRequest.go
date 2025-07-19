// amf/types.go
package amf

import "project/common"

type CreateSessionRequest struct {
	Supi         string          `json:"supi"`
	Gpsi         string          `json:"gpsi"`
	PduSessionID int             `json:"pduSessionId"`
	Dnn          string          `json:"dnn"`
	SNssai       common.SNssai   `json:"sNssai"`
	ServingNHd   string          `json:"servingNHd"`
	AnType       string          `json:"anType"`
}