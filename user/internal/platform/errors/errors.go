package errors

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	// EmailExists indicates a create request with an existing email
	EmailExists = Error("The email address is already associated with another account")
)
