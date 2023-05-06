package errno

type ApiError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

var (
	OK             = ApiError{Code: 0, Message: "OK"}
	ApiErrorSystem = ApiError{Code: 10001, Message: "System Error"}
	ApiErrorIORead = ApiError{Code: 10002, Message: "IO Read Error"}
	ApiErrorJson   = ApiError{Code: 10003, Message: "Json Error"}
)
