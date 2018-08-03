package errors

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	InvalidJWTToken = Error("Invalid JWT token")
)

