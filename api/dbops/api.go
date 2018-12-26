package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"watch/api/defs"
	"watch/api/utils"
)

// User
func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("insert into users (login_name, pwd) values (?, ?)")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	if _, err = stmtIns.Exec(loginName, pwd); err != nil {
		return err
	}

	defer stmtIns.Close()

	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from users where login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	if err = stmtOut.QueryRow(loginName).Scan(&pwd); err != nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmtOut.Close()

	return pwd, nil
}

func DeleteUserCredential(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where login_name=? and pwd=?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	if _, err = stmtDel.Exec(loginName, pwd); err != nil {
		return err
	}

	defer stmtDel.Close()

	return nil
}

// Video
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare("INSERT into video_info (id, author_id, name, display_ctime) values (?,?,?,?)")
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	defer stmtIns.Close()

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {

	stmtOut, err := dbConn.Prepare("select author_id, name, display_ctime from video_info where id=?")
	if err != nil {
		return nil, err
	}

	var aid int
	var name string
	var dist string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dist)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dist}
	return res, nil
}

func DeleteVideo(vid string) error {
	stmtDel, err := dbConn.Prepare("delete from video_info where id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()

	return nil
}
