package services

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/KhanhLinh2810/5G-core/smf/internal/types"
)

var (
	pfcpConnPool = make(map[string]*net.UDPConn)
	pfcpConnLock sync.RWMutex
)

func getOrCreateUDPConn(addr string) (*net.UDPConn, error) {
	pfcpConnLock.RLock()
	conn, exists := pfcpConnPool[addr]
	pfcpConnLock.RUnlock()

	if exists {
		return conn, nil
	}

	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve UPF address: %w", err)
	}

	newConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial UPF: %w", err)
	}

	pfcpConnLock.Lock()
	pfcpConnPool[addr] = newConn
	pfcpConnLock.Unlock()

	return newConn, nil
}

func SendPFCPJsonUDP(msg *types.PFCPMessage, upfAddr string) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal PFCP message: %w", err)
	}

	conn, err := getOrCreateUDPConn(upfAddr)
	if err != nil {
		return fmt.Errorf("get UDP connection failed: %w", err)
	}

	_, err = conn.Write(data)
	if err != nil {
		return fmt.Errorf("failed to send PFCP: %w", err)
	}

	conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return fmt.Errorf("failed to receive PFCP response: %w", err)
	}

	var resp types.PFCPResponse
	if err := json.Unmarshal(buf[:n], &resp); err != nil {
		return fmt.Errorf("failed to parse PFCP response: %w", err)
	}

	return nil
}
