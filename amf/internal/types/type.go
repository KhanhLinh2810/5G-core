package types

type UpdateSession struct {
	Supi         string          `json:"supi"`
	Gpsi         string          `json:"gpsi"`
	PduSessionID int             `json:"pduSessionId"`
	Dnn          string          `json:"dnn"`
	Action       string          `json:"action"`
}

type Session struct {
	Supi         string          `json:"supi"`
	Gpsi         string          `json:"gpsi"`
	PduSessionID int             `json:"pduSessionId"`
	Dnn          string          `json:"dnn"`
}
