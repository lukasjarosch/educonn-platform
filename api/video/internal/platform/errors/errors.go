package errors

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	MissingVideoId = Error("missing video id")
)

