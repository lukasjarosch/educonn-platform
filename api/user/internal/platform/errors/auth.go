package errors

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	InvalidJWTToken = Error("invalid JWT token")
	EmailMissing = Error("email missing")
	PasswordMissing = Error("password missing")
)

