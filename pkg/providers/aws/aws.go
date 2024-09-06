package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AWSClient struct {
	EC2 *ec2.EC2
	S3  *s3.S3
	IAM *iam.IAM
	Session *session.Session
}

func NewAWSClient(accessKey, secretKey, region string) (*AWSClient, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	return &AWSClient{
		EC2:     ec2.New(sess),
		S3:      s3.New(sess),
		IAM:     iam.New(sess),
		Session: sess,
	}, nil
}