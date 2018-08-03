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
	MissingUserId = Error("The request does not contain the user ID")
	MissingEmail = Error("The request is missing an email field")
	MissingPassword = Error("The request is missing the password field")
	InvalidCredentials = Error("The email and/or the password do not match or are incorrect")

	PrivateKeyFileNotFound = Error("The private key could not be found, check the path")
	PublicKeyFileNotFound = Error("The public key could not be found, check the path")

	JwtDecodingFailed = Error("The JWT token could not be decoded. Token is invalid and/or malformed")
)
