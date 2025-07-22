package types

type PFCPMessage struct {
	MessageType int    `json:"messageType"`
	PDNType     string `json:"pdnType"`
	IPAddress   string `json:"ipAddress"`
	SessionID   string `json:"sessionId"`
}

type PFCPResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
