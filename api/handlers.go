package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"watch/api/dbops"
	"watch/api/defs"
	"watch/api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	body := &defs.UserCredential{}

	if err := json.Unmarshal(res, body); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(body.Username, body.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDbError)
		return
	}

	sid := session.GenerateNewSessionId(body.Username)
	signedUp := &defs.SignedUp{Success: true, SessionId: sid}

	if resp, err := json.Marshal(signedUp); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("user_name")
	io.WriteString(w, username)
}
