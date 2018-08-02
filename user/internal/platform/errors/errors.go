package errors

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	// EmailExists indicates a create request with an existing email
	EmailExists = Error("The email address is already associated with another account")
	UserNotFound = Error("No user found")
	MissingIdOrEmail = Error("Missing user_id or email")
	MalformedUserId = Error("The user_id is malformed")
)
