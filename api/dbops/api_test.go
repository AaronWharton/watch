package dbops

import (
	"strconv"
	"testing"
	"time"
)

var tempVid string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

// User test
func TestUserWorkFlow(t *testing.T) {
	t.Run("AddUser", testAddUser)
	t.Run("GetUser", testGetUser)
	t.Run("DelUser", testDeleteUser)
	t.Run("RegetUser", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("aaron", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("aaron")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUserCredential("aaron", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("aaron")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

// video test
func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideo)
	t.Run("GetVideo", testGetVideo)
	t.Run("DeleteVideo", testDeleteVideo)
	t.Run("RegetVideo", testRegetVideo)
}

func testAddVideo(t *testing.T) {
	video, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddNewVideo: %v", err)
	}
	tempVid = video.Id
}

func testGetVideo(t *testing.T) {
	_, err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideo(t *testing.T) {
	err := DeleteVideo(tempVid)
	if err != nil {
		t.Errorf("Error of DeleteVideo: %v", err)
	}
}

func testRegetVideo(t *testing.T) {
	video, err := GetVideoInfo(tempVid)
	if err != nil || video != nil {
		t.Errorf("Error of RegetVideo: %v", err)
	}
}

// comment test
func TestCommentWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddComment", testAddComment)
	t.Run("ListComment", testListComment)
}

func testAddComment(t *testing.T) {

	vid := "12345"
	aid := 1
	content := "I like this video."

	if err := AddNewComment(vid, aid, content); err != nil {
		t.Errorf("Error of AddNewComment: %v", err)
	}
}

func testListComment(t *testing.T) {

	vid := "12345"
	from := 1314521520
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	_, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}
}
