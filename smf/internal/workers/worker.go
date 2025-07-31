package worker

// import (
// 	"fmt"
// 	"log"

// 	"github.com/KhanhLinh2810/5G-core/smf/internal/services"
// 	"github.com/KhanhLinh2810/5G-core/smf/internal/types"
// )

// var JobQueue = make(chan types.Job, 100)

// func StartWorkerPool(num int) {
// 	for i := 0; i < num; i++ {
// 		go func(id int) {
// 			for job := range JobQueue {
// 				switch job.Type {
// 				case types.CreateSession:
// 					req := job.Payload.(types.CreateSessionRequest)
// 					log.Printf("Worker %d handling CREATE for Supi %s", id, req.Supi)
// 					services.CreateSessionSaveInMap(req)

// 				case types.UpdateSession:
// 					req := job.Payload.(services.UpdateSessionRequest)
// 					log.Printf("Worker %d handling UPDATE for Supi %s", id, req.Supi)
// 					services.UpdateSession(req)

// 				case types.ReleaseSession:
// 					req := job.Payload.(services.ReleaseSessionRequest)
// 					log.Printf("Worker %d handling RELEASE for Supi %s", id, req.Supi)
// 					services.ReleaseSession(req)
// 				}
// 			}
// 		}(i)
// 	}
// }
