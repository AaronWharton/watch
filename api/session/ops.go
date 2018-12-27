package session

import "sync"

// session cache
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB() {

}

func GenerateNewSessionId(username string) string {

}

func IsSessionExpired(sid string) (string, error) {

}
