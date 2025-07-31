package controllers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/KhanhLinh2810/5G-core/amf/internal/services"
	"github.com/KhanhLinh2810/5G-core/amf/internal/types"
)

var countResponseN1N2 int32

func UECreateSession(c *gin.Context) {
	// Tạo mock request
	csrJSON := services.MockDataForUERequestHandler()

	// Gửi request tới SMF
	resp, err := services.CreateSession(csrJSON)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": fmt.Sprintf("Failed to send to SMF: %v", err),
		})
		return
	}
	if resp == nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "SMF response is nil"})
		return
	}
	defer resp.Body.Close()

	// Log status
	log.Println("AMF received response from SMF, status:", resp.Status)

	// Trả response về client
	c.JSON(http.StatusOK, gin.H{
		"status":    "Session request processed",
		"smfStatus": resp.Status,
	})
}

func N1N2MessageTransfer(c *gin.Context) {
	var req types.N1N2MessageTransfer

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	atomic.AddInt32(&countResponseN1N2, 1)

	c.JSON(http.StatusOK, gin.H{
		"mess": "success recieve n1n2 message",
	})
}

func MultiUECreateSession(c *gin.Context) {
	numberOfRequest := c.Param("request")
	numRows, err := strconv.Atoi(numberOfRequest)
	if err != nil || numRows <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rowDb, must be integer between 1-400"})
		return
	}
	var countRequest int32
	var wg sync.WaitGroup
	// Tạo mock request
	csrJSON := services.MockDataForUERequestHandler()

	c.JSON(http.StatusOK, gin.H{
		"status":    "Session request processed",
		"smfStatus": "success",
	})

	go func() {
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			countReq := atomic.SwapInt32(&countRequest, 0)
			log.Printf("=================================")
			log.Printf("Requests sent in last 1 second: %d", countReq)
			countRes := atomic.SwapInt32(&countResponseN1N2, 0)
			log.Printf("Response recive in last 1 second: %d", countRes)
		}
	}()

	numberOfRoutine := int(math.Ceil(float64(numRows) / 100.0))
	tickerSendRequest := time.NewTicker(1 * time.Second)
	timeout := time.After(1 * time.Minute)

	for {
		select {
		case <-tickerSendRequest.C:
			for i := 1; i <= numberOfRoutine; i++ {
				wg.Add(1)
				go SendRequestCreateSessionToSMF(&wg, csrJSON, &countRequest)
			}
			wg.Wait()
		case <-timeout:
			log.Println("complete")
			tickerSendRequest.Stop()
			return
		}
	}
}

func SendRequestCreateSessionToSMF(wg *sync.WaitGroup, csrJSON []byte, countRequest *int32) {
	defer wg.Done()
	numberOfRequest := 100
	for i := 0; i < numberOfRequest; i++ {
		// Gửi request tới SMF
		atomic.AddInt32(countRequest, 1)
		resp, err := services.CreateSession(csrJSON)
		if err != nil {
			log.Printf("faild to send smf: ", err)
			return
		}
		if resp == nil {
			log.Printf("SMF response is nil")
			return
		}
		defer resp.Body.Close()

		// Log status
		// log.Println("AMF received response from SMF, status:", resp.Status)
	}
}

// func UpdateUESession(c *gin.Context) {
// 	actions := []string{"DEACTIVE", "HANDOVER", "ACTIVE", "CALL"}
// 	// Gửi request tới SMF
// 	for i := 0; i < 100; i++ {
// 		go func() {
// 			action := actions[rand.Intn(len(actions))]
// 			UpdateSession(action)
// 		}()
// 	}

// 	// Trả response về client
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":    "Session request processed",
// 		"smfStatus": resp.Status,
// 	})
// }

// func ReleaseUESession(c *gin.Context) {
// 	csrJSON := services.MockDataForUpdateUERequestHandler()

// 	// Gửi request tới SMF
// 	resp, err := services.UpdateSession(csrJSON)
// 	if err != nil {
// 		c.JSON(http.StatusBadGateway, gin.H{
// 			"error": fmt.Sprintf("Failed to send to SMF: %v", err),
// 		})
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Trả response về client
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":    "Session request processed",
// 		"smfStatus": resp.Status,
// 	})
// }
