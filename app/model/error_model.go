package model

type Error struct {
    Server bool `json:"server_fault"`
    Message string `json:"message"`
    HttpCode int `json:"code"`
}

func NewServerError(code int, message string) *Error {
    return &Error {
        Server: true,
        Message: message,
        HttpCode: code,
    }
}

func NewClientError(code int, message string) *Error {
    return &Error {
        Server: true,
        Message: message,
        HttpCode: code,
    }
}
