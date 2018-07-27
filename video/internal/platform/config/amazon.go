package config

import "os"

var (
	AwsAccessKey         = os.Getenv("AWS_ACCESS_KEY")
	AwsSecretKey         = os.Getenv("AWS_SECRET_KEY")
	AwsRegion            = os.Getenv("AWS_REGION")
	AwsSqsVideoQueueName = os.Getenv("AWS_VIDEO_QUEUE")
)
