package jwt_handler

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	AuthenticationHeaderMissing = Error("Authentication header missing")
	InvalidAuthenticationScheme = Error("Invalid authentication scheme, use Bearer")
	Base64EncodingError = Error("Base64 encoding error")
	TokenPublicKeyFileNotFound = Error("The public key for the auth tokens could not be found, check the path")
	TokenPrivateKeyFileNotFound = Error("The private key for the auth tokens could not be found, check the path")
	Unauthorized = Error("You are not authorized to perform this action")
	PrivateKeyMissing = Error("The JwtTokenHandler does not have access to the private key and therefore cannot sign a token")
)
