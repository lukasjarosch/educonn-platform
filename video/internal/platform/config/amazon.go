package config

import "os"
import _ "github.com/joho/godotenv/autoload"

var (
	AwsAccessKey         = os.Getenv("AWS_ACCESS_KEY")
	AwsSecretKey         = os.Getenv("AWS_SECRET_KEY")
	AwsRegion            = os.Getenv("AWS_REGION")
	AwsS3VideoBucket     = os.Getenv("AWS_S3_VIDEO_BUCKET")
)
