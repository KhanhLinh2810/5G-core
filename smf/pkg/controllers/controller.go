package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/KhanhLinh2810/5G-core/smf/internal/types"
	"github.com/KhanhLinh2810/5G-core/smf/internal/worker"
)

func CreateSessionSaveInMap(c *gin.Context) {
	var req types.CreateSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	worker.JobQueue <- types.Job{
		Type:    types.CreateSessionJob,
		Payload: req,
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status": "Create session job enqueued",
	})
}

func UpdateSession(c *gin.Context) {
	var req types.Session
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	worker.JobQueue <- types.Job{
		Type:    types.UpdateSessionJob,
		Payload: req,
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status": "Update session job enqueued",
	})
}

func ReleaseSession(c *gin.Context) {
	var req types.ReleaseSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	worker.JobQueue <- types.Job{
		Type:    types.ReleaseSessionJob,
		Payload: req,
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status": "Release session job enqueued",
	})
}
