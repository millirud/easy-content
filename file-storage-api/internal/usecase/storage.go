package usecase

import (
	"context"
	"io"
	"storage-api/internal/entity"
	"storage-api/pkg/requestid"

	"go.uber.org/zap"
)

type (
	storage interface {
		PutObject(
			ctx context.Context,
			reader io.Reader,
			objectSize int64,
			contentType string,
			metadata map[string]string,
		) (*entity.UploadedInfo, error)

		GetObject(
			ctx context.Context,
			bucket string,
			filename string,
		) (*entity.File, func(), error)
	}
)

type storageUseCase struct {
	storage storage
	log     *zap.Logger
}

func NewStorageUseCase(log *zap.Logger, storage storage) *storageUseCase {
	log.Info("NewStorageUseCase")

	return &storageUseCase{
		storage: storage,
		log:     log,
	}
}

func (s *storageUseCase) Upload(ctx context.Context, file entity.File) (*entity.UploadedInfo, error) {
	return s.storage.PutObject(
		ctx,
		file.Reader(),
		file.Size(),
		file.ContentType(),
		map[string]string{
			"basename":   file.Filename(),
			"request-id": requestid.GetRequestId(ctx),
		},
	)
}

func (s *storageUseCase) Get(ctx context.Context, bucket string, filename string) (*entity.File, func(), error) {
	return s.storage.GetObject(ctx, bucket, filename)
}
