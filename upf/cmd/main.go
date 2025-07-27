package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

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

// Redis client
var rdb *redis.Client

func init() {
	// Initialize Redis client
	rdb = redis.NewClient(&redis.Options{
		Addr:     "db:6379", // Redis server address
		Password: "",                // No password
		DB:       0,                // Default DB
	})

	// Test Redis connection
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
}

func main() {
	// Create UDP server
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8805")
	if err != nil {
		fmt.Printf("Failed to resolve UDP address: %v\n", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Printf("Failed to start UDP server: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("UPF server listening on :8805...")

	// Buffer for reading incoming messages
	buffer := make([]byte, 1024)

	for {
		// Read incoming UDP message
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		fmt.Printf("UDP received from %v: %s\n", remoteAddr, string(buffer[:n]))

		if err != nil {
			fmt.Printf("Error reading UDP message: %v\n", err)
			continue
		}

		// Parse PFCP message
		var pfcpMsg PFCPMessage
		if err := json.Unmarshal(buffer[:n], &pfcpMsg); err != nil {
			fmt.Printf("Error parsing PFCP message: %v\n", err)
			sendResponse(conn, remoteAddr, PFCPResponse{
				Status:  "error",
				Message: fmt.Sprintf("Invalid PFCP message: %v", err),
			})
			continue
		}

		// Validate PFCP message
		if pfcpMsg.MessageType != 50 {
			fmt.Printf("Invalid message type: %d\n", pfcpMsg.MessageType)
			sendResponse(conn, remoteAddr, PFCPResponse{
				Status:  "error",
				Message: "Invalid message type, expected 50 (PFCP Session Establishment)",
			})
			continue
		}

		if pfcpMsg.PDNType != "IPv4" && pfcpMsg.PDNType != "IPv6" && pfcpMsg.PDNType != "IPv4v6" {
			fmt.Printf("Invalid PDN type: %s\n", pfcpMsg.PDNType)
			sendResponse(conn, remoteAddr, PFCPResponse{
				Status:  "error",
				Message: "Invalid PDN type",
			})
			continue
		}

		// Store session in Redis
		ctx := context.Background()
		key := fmt.Sprintf("pfcp:session:%s", pfcpMsg.SessionID)
		data := map[string]interface{}{
			"pdnType":   pfcpMsg.PDNType,
			"ipAddress": pfcpMsg.IPAddress,
		}
		if err := rdb.HMSet(ctx, key, data).Err(); err != nil {
			fmt.Printf("Error storing session in Redis: %v\n", err)
			sendResponse(conn, remoteAddr, PFCPResponse{
				Status:  "error",
				Message: fmt.Sprintf("Failed to store session: %v", err),
			})
			continue
		}

		// Send success response
		fmt.Printf("send response ...")
		sendResponse(conn, remoteAddr, PFCPResponse{
			Status:  "success",
			Message: fmt.Sprintf("PFCP Session established for Session ID: %s", pfcpMsg.SessionID),
		})
	}
}

// sendResponse sends a response back to the client
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