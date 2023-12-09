package internal

type APIError struct {
	Messgage string
	Code     int
}

func NewAPIError(message string, code int) APIError {
	return APIError{
		Messgage: message,
		Code:     code,
	}
}
