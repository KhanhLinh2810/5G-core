package types

type N1N2MessageTransfer struct {
	PduSessionID int             `json:"pduSessionId"`
	Dnn          string          `json:"dnn"`
	SNssai       SNssai   `json:"sNssai"`
	AnType       string          `json:"anType"`
}