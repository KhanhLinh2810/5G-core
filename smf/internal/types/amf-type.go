package types

type SNssai struct {
	Sst int    `json:"sst"`
	Sd  string `json:"sd"`
}

type CreateSessionRequest struct {
	Supi         string          `json:"supi"`
	Gpsi         string          `json:"gpsi"`
	PduSessionID int             `json:"pduSessionId"`
	Dnn          string          `json:"dnn"`
	SNssai       SNssai   `json:"sNssai"`
	ServingNHd   string          `json:"servingNHd"`
	AnType       string          `json:"anType"`
}