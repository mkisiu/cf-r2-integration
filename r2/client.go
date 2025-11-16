package r2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Config - configuration params - connection to R2
type Config struct {
	Endpoint string
	AccessKeyID string
	SecretAccessKey string
	BucketName string
}

// Client - client to connect to R2
type Client struct {
	s3 *s3.Client
	bucket string
}

// NewClient - creates an instance of the wrapper around the S3 client for ease of use
func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	// check if anything is missing to fail early
	if cfg.Endpoint == "" || cfg.AccessKeyID == "" || cfg.SecretAccessKey == "" || cfg.BucketName == "" {
		return nil, fmt.Errorf("r2: incomplete configuration")
	}

	// create AWS S3 config (we are using underlying S3 client)
	awsCfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("auto"),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.AccessKeyID,
				cfg.SecretAccessKey,
				"", // this is for session tokens - we're not using it here
			),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("r2: aws config: %w", err)
	}

	// having above data ready, instantiate the S3 client
	s3client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(cfg.Endpoint)
		o.UsePathStyle = true
	})

	// return what we promissed (our internal Client (wrapper around S3), and an error)
	return &Client{
		s3: s3client,
		bucket: cfg.BucketName,
	}, nil
}

