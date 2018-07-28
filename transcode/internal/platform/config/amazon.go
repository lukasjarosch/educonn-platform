package config

import "os"

var (
	AwsAccessKey         = os.Getenv("AWS_ACCESS_KEY")
	AwsSecretKey         = os.Getenv("AWS_SECRET_KEY")
	AwsRegion            = os.Getenv("AWS_REGION")
	AwsSqsVideoQueueName = os.Getenv("AWS_VIDEO_QUEUE")

	AwsTranscodePresetId     = os.Getenv("AWS_TRANSCODE_PRESET")
	AwsTranscodePipelineId   = os.Getenv("AWS_TRANSCODE_PIPELINE")
	AwsTranscodeOutputPrefix = os.Getenv("AWS_TRANSCODE_OUTPUT_PREFIX")
)

