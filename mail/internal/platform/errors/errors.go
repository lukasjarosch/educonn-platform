package errors

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	InvalidEmail = Error("The email address is invalid")
	SqsNoMessages = Error("No SQS messages retrieved")
)
