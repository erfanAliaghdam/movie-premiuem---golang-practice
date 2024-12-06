package utils

import (
	"bytes"
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

	// Specify the destination key in the bucket
	destinationKey := "golang/" + fileName

	// PreSign the GET object request
	preSignedUrl, preSignedUrlErr := preSignClient.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(appConfig.BucketName),
		Key:    aws.String(destinationKey),
	}, func(options *s3.PresignOptions) {
		options.Expires = time.Minute * 15
	})

	if preSignedUrlErr != nil {
		return "", preSignedUrlErr
	}

	return preSignedUrl.URL, nil

}

func UploadFileToBucket(fileContent *bytes.Reader, fileName string) error {
	appConfig := core.LoadConfig()

	cfg, err := bucketConfig.LoadDefaultConfig(context.TODO(), bucketConfig.WithRegion("us-west-2"))
	if err != nil {
		return err
	}

	// Define AWS credentials and bucket information
	awsAccessKeyID := appConfig.BucketAccessKey
	awsSecretAccessKey := appConfig.SecretKey
	endpoint := appConfig.BucketEndpoint
	bucketName := appConfig.BucketName

	// Initialize S3 client with custom configuration
	cfg.Credentials = aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     awsAccessKeyID,
			SecretAccessKey: awsSecretAccessKey,
		}, nil
	})

	cfg.BaseEndpoint = aws.String(endpoint)

	client := s3.NewFromConfig(cfg)

	// Specify the destination key in the bucket
	destinationKey := "golang/" + fileName

	// Use the S3 client to upload the file
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(destinationKey),
		Body:   fileContent,
	})

	return err
}
