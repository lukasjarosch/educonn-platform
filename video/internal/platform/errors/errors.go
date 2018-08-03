package errors

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	RawVideoFileS3NotFound = Error("Video file was not found in S3")
	RawVideoAlreadyExists = Error("The raw video key is already associated with another video")
	MissingVideoId = Error("The video id is missing")
	MissingUserId = Error("The user id is missing")
	VideoNotFound = Error("Video not found")
	InvalidVideoId = Error("The provided video id is malformed")
)

