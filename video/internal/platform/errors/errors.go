package errors

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	RawVideoFileS3NotFound = Error("Video file was not found in S3")
)

