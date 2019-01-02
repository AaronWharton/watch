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
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC: 400, Error: Err{Error: "Bad request!", ErrorCode: "001"}}
	ErrorUserUnauthorized       = ErrResponse{HttpSC: 401, Error: Err{Error: "User authentication failed!", ErrorCode: "002"}}
	ErrorDbError                = ErrResponse{HttpSC: 500, Error: Err{Error: "DB operations failed", ErrorCode: "003"}}
	ErrorInternalFaults         = ErrResponse{HttpSC: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
)
