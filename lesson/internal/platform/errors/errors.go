package errors

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	UnknownLessonType = Error("unknown lesson type")
	MissingVideoId = Error("missing videoId")
	MissingLessonId = Error("missing lessonId")
	MissingType = Error("missing type")
	MissingUserId = Error("missing userId")
	MissingLessonName = Error("missing lesson name")
	MongoCreateFailed = Error("failed to create mongodb document")
	MalformedId = Error("malformed bson id")
	MissingUserIdHeader = Error("missing x-user-id in context")
)

