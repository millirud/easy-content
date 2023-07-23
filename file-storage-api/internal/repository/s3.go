package repository

import (
	"context"
	"io"
	"storage-api/internal/entity"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

type s3 struct {
	log    *zap.Logger
	client *minio.Client
	bucket string
}

func NewS3(
	log *zap.Logger,

	endpoint string,
	accessKeyID string,
	secretAccessKey string,
	useSSL bool,
	bucket string,
) (*s3, error) {
	// Initialize minio client object.
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Error("new minio client error", zap.Error(err))

		return nil, err
	}

	return &s3{
		log:    log,
		client: client,
		bucket: bucket,
	}, nil

}

func (s *s3) getClient() *minio.Client {
	return s.client
}

func (s *s3) generateFileName() string {
	id := uuid.New()
	return id.String()
}

func (s *s3) PutObject(
	ctx context.Context,
	reader io.Reader,
	objectSize int64,
	contentType string,
	metadata map[string]string,
) (*entity.UploadedInfo, error) {
	filename := s.generateFileName()

	_, err := s.getClient().
		PutObject(
			ctx,
			s.bucket,
			filename,
			reader,
			objectSize,
			minio.PutObjectOptions{
				ContentType:  contentType,
				UserMetadata: metadata,
			},
		)

	if err != nil {
		s.log.Error("s3.PutObject", zap.Error(err))
		return nil, err
	}

	return entity.NewUploadedInfo(filename, s.bucket), nil
}

func (s *s3) GetObject(
	ctx context.Context,
	bucket string,
	filename string,
) (*entity.File, func(), error) {
	obj, err := s.getClient().GetObject(ctx, bucket, filename, minio.GetObjectOptions{})

	if err != nil {
		return nil, nil, err
	}

	stat, err := obj.Stat()

	if err != nil {
		return nil, nil, err
	}

	return entity.NewFile(
			obj,
			stat.Size,
			stat.ContentType,
			filename,
		),
		func() {
			obj.Close()
		},
		nil
}
