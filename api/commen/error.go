package commen

//定义一些通用的respose错误信息

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErroResponse struct {
	HttpSC int
	Error  Err
}

// 实例化一下常见的错误类型

var (
	ErroRequestBodyParseFailed = ErroResponse{HttpSC: 400, Error: Err{Error: "request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser           = ErroResponse{HttpSC: 401, Error: Err{Error: "User authentication failed.", ErrorCode: "002"}}
	ErrorDBError               = ErroResponse{HttpSC: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"}}
	ErrorInternalFaults        = ErroResponse{HttpSC: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
)
