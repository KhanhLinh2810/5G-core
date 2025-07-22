package types

type N1N2TransferRequest struct {
	PduSessionID int      `json:"pduSessionId"`
	SNssai       SNssai   `json:"sNssai"`
	Dnn          string   `json:"dnn"`
}