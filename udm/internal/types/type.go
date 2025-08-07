package types

type JobType string

const (
	GetSDMDetail      JobType = "GET_SDM_DETAIL"
)

type Job struct {
	ID      int
	Type    JobType
	Payload interface{}
	ResultChan chan any
}

type GetSDMDetailType struct {
	Supi string
}
