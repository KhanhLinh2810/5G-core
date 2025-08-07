package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/KhanhLinh2810/5G-core/udm/internal/types"
	"github.com/KhanhLinh2810/5G-core/udm/internal/workers"
)

func GetSDMDetail(c *gin.Context) {
	supi := c.Param("imsi")
	if supi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing supi parameter"})
		return
	}

	resultChan := make(chan any)
	job := types.Job{
		Type:       types.GetSDMDetail,
		Payload:    types.GetSDMDetailType{Supi: supi},
		ResultChan: resultChan,
	}

	workers.JobQueue <- job

	result := <-resultChan
	response, ok := result.(map[string]any)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response from worker"})
		return
	}

	status := response["status"].(int)
	data := response["data"]

	c.JSON(status, data)
}
