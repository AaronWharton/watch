package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
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
