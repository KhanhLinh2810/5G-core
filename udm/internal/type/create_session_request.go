package types

import "common"

type TypeCreateSessionRequest struct {
	supi         string          `json:"supi"`
	gpsi         string          `json:"gpsi"`
	pduSessionID int             `json:"pduSessionId"`
	dnn          string          `json:"dnn"`
	sNssai       common.SNssai   `json:"sNssai"`
	servingNHd   string          `json:"servingNHd"`
	anType       string          `json:"anType"`
}