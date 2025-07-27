package models

import (
	"encoding/json"
	"fmt"

	"github.com/KhanhLinh2810/5G-core/smf/pkg/config"
	"github.com/KhanhLinh2810/5G-core/smf/internal/types"
)

func sessionKey(supi string) string {
	return fmt.Sprintf("smf:session:%s", supi)
}

func SaveSession(s *types.CreateSessionRequest) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return config.Rdb.Set(config.Ctx, sessionKey(s.Supi), data, 0).Err()
}