package errno

type ApiError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

var (
	OK             = ApiError{Code: 0, Message: "OK"}
	ApiErrorSystem = ApiError{Code: 10001, Message: "System Error"}
)
