package main

import (
	"net/http"
	"watch/api/defs"
	"watch/api/session"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_USERNAME = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	username, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	r.Header.Add(HEADER_FIELD_USERNAME, username)
	return true
}

func validateUser(w http.ResponseWriter, r *http.Request) bool {
	username := r.Header.Get(HEADER_FIELD_USERNAME)
	if len(username) == 0 {
		sendErrorResponse(w, defs.ERROR_UESER_UNAUTHORIZED)
		return false
	}

	return true
}
