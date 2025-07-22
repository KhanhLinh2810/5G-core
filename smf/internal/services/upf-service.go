// smf/internal/services/pfcp.go
package services

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/KhanhLinh2810/5G-core/smf/internal/types"
)

func SendPFCPJsonUDP(msg *types.PFCPMessage, upfAddr string) (*types.PFCPResponse, error) {
	// Serialize PFCPMessage to JSON
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal PFCP message: %w", err)
	}

	// Create UDP connection
	conn, err := net.Dial("udp", upfAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial UPF: %w", err)
	}
	defer conn.Close()

	// Send message
	_, err = conn.Write(data)
	if err != nil {
		return nil, fmt.Errorf("failed to send PFCP: %w", err)
	}

	// Set timeout and wait for response
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to receive PFCP response: %w", err)
	}

	var resp types.PFCPResponse
	if err := json.Unmarshal(buf[:n], &resp); err != nil {
		return nil, fmt.Errorf("failed to parse PFCP response: %w", err)
	}

	return &resp, nil
}
