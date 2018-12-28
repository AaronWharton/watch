package defs

// Model

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

type Comment struct {
	Id       string
	VideoId  string
	UserName string
	Content  string
}

type SimpleSession struct {
	UserName string
	TTL      int64
}
