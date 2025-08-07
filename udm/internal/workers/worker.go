package workers

import (
	"github.com/KhanhLinh2810/5G-core/udm/internal/services"
	"github.com/KhanhLinh2810/5G-core/udm/internal/types"
)

var JobQueue = make(chan types.Job, 100)

func StartWorkerPool(num int) {
	for i := 0; i < num; i++ {
		go func(id int) {
			for job := range JobQueue {
				switch job.Type {
				case types.GetSDMDetail:
					req := job.Payload.(types.GetSDMDetailType)
					services.GetSDMDetail(req, job.ResultChan)
				}
			}
		}(i)
	}
}
