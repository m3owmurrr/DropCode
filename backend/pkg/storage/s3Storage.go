package storage

import (
	"context"
	"errors"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/m3owmurrr/dropcode/backend/pkg/config"
)

type S3Storage struct {
	cl *s3.Client
}

func NewS3Storage(cfg config.S3Config) (*S3Storage, error) {
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{
				URL:           cfg.Endpoint,
				SigningRegion: cfg.Region,
			},
			nil
	})

	ctx := context.Background()
	awsCfg, err := awsConfig.LoadDefaultConfig(ctx,
		awsConfig.WithRegion(cfg.Region),
		awsConfig.WithEndpointResolver(customResolver),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AccessKey, cfg.SecretKey, "",
		)),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	return &S3Storage{cl: client}, nil
}

func (s3s *S3Storage) Put(ctx context.Context, bucket string, key string, data io.Reader) error {
	_, err := s3s.cl.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        data,
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		log.Printf("can't put object %s/%s: %v\n", bucket, key, err)
		return err
	}

	return nil
}

func (s3s *S3Storage) Get(ctx context.Context, bucket string, key string) (io.Reader, error) {
	data, err := s3s.cl.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		var noKey *types.NoSuchKey
		if errors.As(err, &noKey) {
			log.Printf("can't get object %s/%s: no such key exists\n", bucket, key)
		} else {
			log.Printf("can't get object %s/%s: %v", bucket, key, err)
		}
		return nil, err
	}

	return data.Body, nil
}
