package amazon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elastictranscoder"
	"fmt"
	"github.com/lukasjarosch/educonn-platform/transcode/internal/platform/config"
	"github.com/rs/xid"
)

type ElasticTranscoderClient struct {
	svc *elastictranscoder.ElasticTranscoder
}

func NewElasticTranscoderClient(accessKey string, secretKey string, region string) (*ElasticTranscoderClient, error) {
	creds := credentials.NewStaticCredentials(accessKey, secretKey, "")
	_, err := creds.Get()
	if err != nil {
		return nil, err
	}

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	et := &ElasticTranscoderClient{
		svc: elastictranscoder.New(sess),
	}

	return et, nil
}

func (e *ElasticTranscoderClient) CreateJob(inputKey string) (*elastictranscoder.CreateJobResponse, error){
	presetId := config.AwsTranscodePresetId
	pipelineId := config.AwsTranscodePipelineId
	outputKeyPrefix := config.AwsTranscodeOutputPrefix

	// new unique filename
	id := xid.New()
	filename := fmt.Sprintf("%s.mp4", id.String())

	// create job
	resp, err := e.svc.CreateJob(&elastictranscoder.CreateJobInput{
		Input: &elastictranscoder.JobInput{
			Key: aws.String(inputKey),
		},
		OutputKeyPrefix: aws.String(outputKeyPrefix),
		Outputs: []*elastictranscoder.CreateJobOutput{
			{
				PresetId: aws.String(presetId),
				Key:      aws.String(filename),
			},
		},
		PipelineId: aws.String(pipelineId),
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

