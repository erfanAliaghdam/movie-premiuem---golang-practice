package utils

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	bucketConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"movie_premiuem/core"
	"time"
)

func GeneratePreSignedURL(fileName string) (string, error) {
	appConfig := core.LoadConfig()

	// initialize bucket
	bucketCfg, bucketCfgErr := bucketConfig.LoadDefaultConfig(context.TODO(), bucketConfig.WithRegion("us-west-2"))
	if bucketCfgErr != nil {
		return "", bucketCfgErr
	}

	bucketCfg.Credentials = aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     appConfig.BucketAccessKey,
			SecretAccessKey: appConfig.BucketSecretKey,
		}, nil
	})
	bucketCfg.BaseEndpoint = aws.String(appConfig.BucketEndpoint)

	// Initialize S3 pre sign client
	s3client := s3.NewFromConfig(bucketCfg)
	preSignClient := s3.NewPresignClient(s3client)

	// PreSign the GET object request
	preSignedUrl, preSignedUrlErr := preSignClient.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(appConfig.BucketName),
		Key:    aws.String(fileName),
	}, func(options *s3.PresignOptions) {
		options.Expires = time.Minute * 15
	})

	if preSignedUrlErr != nil {
		return "", preSignedUrlErr
	}

	return preSignedUrl.URL, nil

}
