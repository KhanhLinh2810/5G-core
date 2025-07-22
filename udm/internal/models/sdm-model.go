package models

import (
	"encoding/json"
	"fmt"

	"github.com/KhanhLinh2810/5G-core/udm/pkg/config"
	"github.com/KhanhLinh2810/5G-core/udm/internal/types"
)

func sessionKey(supi string) string {
	return fmt.Sprintf("smd:session:%s", supi)
}

func CreateSession(s *types.SDMData) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return config.Rdb.Set(config.Ctx, sessionKey(s.Supi), data, 0).Err()
}

func GetSessionBySupi(supi string) (*types.SDMData, error) {
	data, err := config.Rdb.Get(config.Ctx, sessionKey(supi)).Result()
	if err != nil {
		return nil, err
	}
	var s types.SDMData
	if err := json.Unmarshal([]byte(data), &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func GetSessionList() ([]*types.SDMData, error) {
	var sessions []*types.SDMData

	iter := config.Rdb.Scan(config.Ctx, 0, "smd:session:*", 0).Iterator()
	for iter.Next(config.Ctx) {
		key := iter.Val()
		val, err := config.Rdb.Get(config.Ctx, key).Result()
		if err != nil {
			return nil, fmt.Errorf("error reading key %s: %w", key, err)
		}
		var sdm types.SDMData
		if err := json.Unmarshal([]byte(val), &sdm); err != nil {
			return nil, fmt.Errorf("error unmarshalling key %s: %w", key, err)
		}
		sessions = append(sessions, &sdm)
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return sessions, nil
}