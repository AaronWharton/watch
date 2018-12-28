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
	ERROR_REAUEST_BODY_PARSE_FAILED = ErrResponse{HttpSC: 400, Error: Err{Error: "Bad request!", ErrorCode: "001"}}
	ERROR_UESER_UNAUTHORIZED        = ErrResponse{HttpSC: 401, Error: Err{Error: "User authentication failed!", ErrorCode: "002"}}
	ERROR_DB_ERROR                  = ErrResponse{HttpSC: 500, Error: Err{Error: "DB operations failed", ErrorCode: "003"}}
	ERROR_INTERNAL_FAULTS           = ErrResponse{HttpSC: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
)
