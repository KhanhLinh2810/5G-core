package controllers

import (
	"strconv"
	"net/http"
	"math/rand"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
	
	"github.com/KhanhLinh2810/5G-core/udm/internal/models"
	"github.com/KhanhLinh2810/5G-core/udm/internal/types"
)

func GetSDMDetail(c *gin.Context) {
	supi := c.Param("imsi")
	if supi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing supi parameter"})
		return
	}

	session, err := models.GetSessionBySupi(supi)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "No session found for SUPI",
			"supi":    supi,
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"supi":     session.Supi,
		"dnn":      session.Dnn,
		"stt":  	session.Sst,
		"sd":  		session.Sd,
	})
}

func CreateSession(c *gin.Context) {
	rowStr := c.Param("rowDb")
	numRows, err := strconv.Atoi(rowStr)
	if err != nil || numRows <= 0 || numRows > 400 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rowDb, must be integer between 1-400"})
		return
	}

	// Các giá trị giả định có thể có
	dnnList := []string{"internet", "ims", "enterprise", "iot"}
	sdList := []string{"010203", "112233", "445566", "778899"}
	sstList := []int{1, 2, 3}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numRows; i++ {
		sdm := &types.SDMData{
			Supi: fmt.Sprintf("imsi-20893%09d", i+1),
			Dnn:  dnnList[rand.Intn(len(dnnList))],
			Sst:  sstList[rand.Intn(len(sstList))],
			Sd:   sdList[rand.Intn(len(sdList))],
		}
		if err := models.CreateSession(sdm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed at index %d: %v", i, err)})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Successfully created %d sessions", numRows)})
}

func GetSessions(c *gin.Context) {
	sessions, err := models.GetSessionList()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"sessions": sessions})
}