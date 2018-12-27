package session

import (
	"sync"
	"time"
	"watch/api/dbops"
	"watch/api/defs"
	"watch/api/utils"
)

// session cache
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB() {
	m, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}

	m.Range(func(key, value interface{}) bool {
		simpleSession := value.(*defs.SimpleSession)
		sessionMap.Store(key, simpleSession)
		return true
	})
}

func GenerateNewSessionId(username string) string {
	id, _ := utils.NewUUID()
	ctime := nowInMilliSecond()
	ttl := ctime * 30 * 60 * 1000

	simpleSession := &defs.SimpleSession{UserName: username, TTL: ttl}
	sessionMap.Store(id, simpleSession)
	_ = dbops.InsertSession(id, ttl, username)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	simpleSession, ok := sessionMap.Load(sid)
	if ok {
		ctime := nowInMilliSecond()
		if simpleSession.(*defs.SimpleSession).TTL < ctime {
			deleteExpiredSession(sid)
			return "", true
		}

		return simpleSession.(defs.SimpleSession).UserName, false
	}

	return "", true
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	_ = dbops.DeleteSession(sid)
}

func nowInMilliSecond() int64 {
	return time.Now().UnixNano() / 1000000
}
