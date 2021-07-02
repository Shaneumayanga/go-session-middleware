package gosession

import (
	"time"
)

func RunCleanUp(s *Store, sessionId string) {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			s.DeleteExpiredSession(sessionId)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
