package types

type ReleaseSessionRequest struct {
	Supi         string `json:"supi"`
	Gpsi         string `json:"gpsi"`
	PduSessionID int    `json:"pduSessionId"`
	Dnn          string `json:"dnn"`
}

type Session struct {
	Supi         string `json:"supi"`
	Gpsi         string `json:"gpsi"`
	PduSessionID int    `json:"pduSessionId"`
	Dnn          string `json:"dnn"`
	Action       string `json:"action"`
}

type SessionStore struct {
	mu       sync.RWMutex
	sessions map[string]Session
}

type JobType string

const (
	UpdateSMF      JobType = "UPDATE_SMF"
	ProcessMessage JobType = "PROCESS_MESSAGE"
	SendEmail      JobType = "SEND_EMAIL"
)

type Job struct {
	ID      int
	Type    JobType
	Payload interface{}
}


