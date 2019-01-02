package main

import (
	"net/http"
	"watch/api/defs"
	"watch/api/session"
)

var HeaderFieldSession = "X-Session-Id"
var HeaderFieldUsername = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HeaderFieldSession)
	if len(sid) == 0 {
		return false
	}

	username, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	r.Header.Add(HeaderFieldUsername, username)
	return true
}

func validateUser(w http.ResponseWriter, r *http.Request) bool {
	username := r.Header.Get(HeaderFieldUsername)
	if len(username) == 0 {
		sendErrorResponse(w, defs.ErrorUserUnauthorized)
		return false
	}

	return true
}
