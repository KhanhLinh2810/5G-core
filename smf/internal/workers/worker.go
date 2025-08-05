package workers

import (
	// "fmt"
	// "log"

	"github.com/KhanhLinh2810/5G-core/smf/internal/services"
	"github.com/KhanhLinh2810/5G-core/smf/internal/types"
)

var JobQueue = make(chan types.Job, 100)

func StartWorkerPool(num int) {
	for i := 0; i < num; i++ {
		go func(id int) {
			for job := range JobQueue {
				switch job.Type {
				case types.CreateSession:
					req := job.Payload.(types.CreateSessionRequest)
					services.CreateSession(req, job.ResultChan)
				}
			}
		}(i)
	}
}
