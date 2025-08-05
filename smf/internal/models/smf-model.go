package models

import (
	"encoding/json"
	"fmt"

	"github.com/KhanhLinh2810/5G-core/smf/internal/types"
	"github.com/KhanhLinh2810/5G-core/smf/pkg/config"
)

// map

// var GlobalSessionStore = &types.SessionStore{
// 	sessions: make(map[string]Session),
// }

// func (s *types.SessionStore) SaveSessionInMap(session Session) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	s.sessions[session.Supi] = session
// }

// func (s *SessionStore) Get(supi string) (Session, bool) {
// 	s.mu.RLock()
// 	defer s.mu.RUnlock()
// 	session, ok := s.sessions[supi]
// 	return session, ok
// }

// func (s *SessionStore) Delete(supi string) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	delete(s.sessions, supi)
// }

// redis

func sessionKey(supi string) string {
	return fmt.Sprintf("smf:session:%s", supi)
}

func SaveSession(s *types.CreateSessionRequest) error {
	action := "ACTIVE"
	session := types.Session{
		Supi:         s.Supi,
		Gpsi:         s.Gpsi,
		PduSessionID: s.PduSessionID,
		Dnn:          s.Dnn,
		Action:       action,
	}

	data, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return config.Rdb.Set(config.Ctx, sessionKey(s.Supi), data, 0).Err()
}

// // fake data
// func SeedFakeSessions() {
// 	GlobalSessionStore.Save(types.Session{
// 		Supi:         "imsi-452040000000001",
// 		Gpsi:         "msisdn-84900000001",
// 		PduSessionID: 1,
// 		Dnn:          "v-internet",
// 		Action:       "ACTIVE",
// 	})
// 	GlobalSessionStore.Save(types.Session{
// 		Supi:         "imsi-452040000000002",
// 		Gpsi:         "msisdn-84900000002",
// 		PduSessionID: 2,
// 		Dnn:          "v-internet",
// 		Action:       "ACTIVE",
// 	})
// 	GlobalSessionStore.Save(types.Session{
// 		Supi:         "imsi-452040000000003",
// 		Gpsi:         "msisdn-84900000003",
// 		PduSessionID: 3,
// 		Dnn:          "v-internet",
// 		Action:       "ACTIVE",
// 	})
// }
