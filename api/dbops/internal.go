package dbops

import (
	"database/sql"
	"strconv"
	"sync"
	"watch/api/defs"
)

func InsertSession(sid string, ttl int64, username string) error {

	ttlStr := strconv.FormatInt(ttl, 10)

	stmtIns, err := dbConn.Prepare("INSERT into sessions (session_id, TTL, login_name) values (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(sid, ttlStr, username)
	if err != nil {
		return err
	}

	defer stmtIns.Close()

	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {

	simpleSession := &defs.SimpleSession{}

	stmtOut, err := dbConn.Prepare("select TTL, login_name from sessions where session_id=?")
	if err != nil {
		return nil, err
	}

	var username string
	var ttl string

	err = stmtOut.QueryRow(sid).Scan(&ttl, &username)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if ttlInt, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		simpleSession.UserName = username
		simpleSession.TTL = ttlInt
	} else {
		return nil, err
	}

	defer stmtOut.Close()

	return simpleSession, nil
}

func RetrieveAllSessions() (*sync.Map, error) {

	m := &sync.Map{}

	stmtOut, err := dbConn.Prepare("SELECT * from sessions")
	if err != nil {
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttl string
		var loginName string

		if err1 := rows.Scan(&id, &ttl, &loginName); err1 != nil {
			break
		}

		if ttlInt, err2 := strconv.ParseInt(ttl, 10, 64); err2 == nil {
			simpleSession := &defs.SimpleSession{UserName: loginName, TTL: ttlInt}
			m.Store(id, simpleSession)
		}
	}

	return m, nil
}

func DeleteSession(sid string) error {
	stmtDel, err := dbConn.Prepare("DELETE * from sessions where session_id=?")
	if err != nil {
		return err
	}

	if _, err = stmtDel.Exec(sid); err != nil {
		return err
	}

	defer stmtDel.Close()

	return nil
}
