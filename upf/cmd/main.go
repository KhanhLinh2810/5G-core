package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	// "sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// PFCPMessage represents the structure of a PFCP Session Establishment message
type PFCPMessage struct {
	MessageType int    `json:"messageType"`
	PDNType     string `json:"pdnType"`
	IPAddress   string `json:"ipAddress"`
	SessionID   string `json:"sessionId"`
}

// PFCPResponse represents the structure of the PFCP response
type PFCPResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type PFCPJob struct {
	Data       []byte
	RemoteAddr *net.UDPAddr
}

var rdb *redis.Client

// Job queue
var JobQueue = make(chan PFCPJob, 10000)

func main() {
	// Initialize Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:            "db:6379",
		Password:        "",
		DB:              0,
		PoolSize:        4000,
		MinIdleConns:    400,
		PoolTimeout:     30 * time.Second,
		ConnMaxIdleTime: 5 * time.Minute,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	// Start UDP server
	addr, err := net.ResolveUDPAddr("udp", ":8805")
	if err != nil {
		panic(fmt.Sprintf("Failed to resolve UDP address: %v", err))
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(fmt.Sprintf("Failed to start UDP server: %v", err))
	}
	defer conn.Close()
	fmt.Println("UPF server listening on :8805...")

	StartWorkerPool(conn, 100) 

	// Read loop
	buffer := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading UDP messages: %v\n", err)
			continue
		}

		// Push job to queue
		dataCopy := make([]byte, n)
		copy(dataCopy, buffer[:n])
		job := PFCPJob{
			Data:       dataCopy,
			RemoteAddr: remoteAddr,
		}
		JobQueue <- job
	}
}

// StartWorkerPool launches multiple workers
func StartWorkerPool(conn *net.UDPConn, numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			for job := range JobQueue {
				handlePFCPMessage(conn, job.Data, job.RemoteAddr)
			}
		}(i)
	}
}

// handlePFCPMessage processes a single PFCP message
func handlePFCPMessage(conn *net.UDPConn, data []byte, remoteAddr *net.UDPAddr) {
	var pfcpMsg PFCPMessage
	if err := json.Unmarshal(data, &pfcpMsg); err != nil {
		sendResponse(conn, remoteAddr, PFCPResponse{
			Status:  "error",
			Message: fmt.Sprintf("Invalid PFCP messages: %v", err),
		})
		return
	}

	if pfcpMsg.MessageType != 50 {
		sendResponse(conn, remoteAddr, PFCPResponse{
			Status:  "error",
			Message: "Invalid message type, expected 50 (PFCP Session Establishment)",
		})
		return
	}

	if pfcpMsg.PDNType != "IPv4" && pfcpMsg.PDNType != "IPv6" && pfcpMsg.PDNType != "IPv4v6" {
		sendResponse(conn, remoteAddr, PFCPResponse{
			Status:  "error",
			Message: "Invalid PDN type",
		})
		return
	}

	key := fmt.Sprintf("pfcp:session:%s", pfcpMsg.SessionID)
	dataMap := map[string]interface{}{
		"pdnType":   pfcpMsg.PDNType,
		"ipAddress": pfcpMsg.IPAddress,
	}

	if err := rdb.HMSet(context.Background(), key, dataMap).Err(); err != nil {
		sendResponse(conn, remoteAddr, PFCPResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to store session in Redis: %v", err),
		})
		return
	}

	sendResponse(conn, remoteAddr, PFCPResponse{
		Status:  "success",
		Message: fmt.Sprintf("PFCP Session established for Session ID: %s", pfcpMsg.SessionID),
	})
}

// sendResponse sends a PFCP response via UDP
func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, response PFCPResponse) {
	respBytes, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error marshaling response: %v\n", err)
		return
	}
	_, err = conn.WriteToUDP(respBytes, addr)
	if err != nil {
		fmt.Printf("Error sending response: %v\n", err)
	}
}
