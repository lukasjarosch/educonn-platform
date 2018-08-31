package amazon

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"time"
	"context"
	"github.com/opentracing/opentracing-go"
)

type S3Bucket struct {
	Bucket    string
	Region    string
	accessKey string
	secretKey string
	session   *s3.S3
}

// New creates a new bucket
func NewS3Bucket(bucket string, region string, accessKey string, secretKey string) (*S3Bucket, error) {
	var b S3Bucket

	b.Bucket = bucket
	b.Region = region
	b.accessKey = accessKey
	b.secretKey = secretKey

	creds := credentials.NewStaticCredentials(accessKey, secretKey, "")
	_, err := creds.Get()
	if err != nil {
		return nil, err
	}

	sess, _ := session.NewSession(&aws.Config{
		Region:   aws.String(b.Region),
	})

	b.session = s3.New(sess)

	return &b, nil
}

// CheckFileExists tries to retrieve the metadata of the file. If no metadata is found, the file does not exist
// If this method does not return an error, the file does exist. A bucket must be set here, so if you plan on using
// the default value, do it on your own.
func (b *S3Bucket) CheckFileExists(ctx context.Context, file string, bucket string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "S3Bucket.CheckFileExists")
	defer span.Finish()

	_, err := b.session.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(file),
	})

	if err != nil {
		return err
	}

	return  nil
}

// Upload uploads a given buffer to S3. This is useful for the client
func (b *S3Bucket) Upload(filePath string, buffer []byte, size int64) (*s3.PutObjectOutput, error) {
	output, err := b.session.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(b.Bucket),
		Key:           aws.String(filePath),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(http.DetectContentType(buffer)),
	})

	if err != nil {
		return nil, err
	}

	return output, nil
}

func (b *S3Bucket) GetSignedResourceURL(fileKey string) (string, error) {
	params := &s3.GetObjectInput{
		Bucket: aws.String(b.Bucket),
		Key: aws.String(fileKey),
	}

	req, _ := b.session.GetObjectRequest(params)
	url, err := req.Presign(15 * time.Minute)	 // TODO: move to config
	if err != nil {
		return "", err
	}
	return url, nil
}