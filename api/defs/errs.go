package defs

// NOTE:Error code is not http status code!
type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorBadRequest   = ErrResponse{HttpSC: 400, Error: Err{Error: "Bad request!", ErrorCode: "001"}}
	ErrorUnauthorized = ErrResponse{HttpSC: 401, Error: Err{Error: "User authentication failed!", ErrorCode: "002 "}}
)
