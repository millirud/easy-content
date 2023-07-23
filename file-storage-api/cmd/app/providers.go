package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"

	"storage-api/internal/entity"
	httphandler "storage-api/internal/http/handler"
	"storage-api/internal/repository"
	"storage-api/internal/usecase"

	"storage-api/config"

	"storage-api/internal/http/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(lc fx.Lifecycle, gin *gin.Engine, log *zap.Logger) *http.Server {
	srv := &http.Server{
		Handler: gin,
		Addr:    ":8080",
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)

			if err != nil {
				return err
			}

			log.Info(fmt.Sprintf("Starting HTTP server at %s", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("shutting down http server")

			return srv.Shutdown(ctx)
		},
	})
	return srv
}

type StorageHttpHandler interface {
	Get(*gin.Context)
	Upload(*gin.Context)
}

func NewGin(
	storageHandler StorageHttpHandler,
) *gin.Engine {
	handler := gin.New()

	handler.Use(middleware.NewLogger())
	handler.Use(middleware.NewRequestidMiddleware())

	handler.GET("/v1/view/:bucket/:filename", storageHandler.Get)
	handler.POST("/v1/upload", storageHandler.Upload)

	return handler
}

func NewStorageHandler(
	log *zap.Logger,
	useCase StorageUseCase,
) StorageHttpHandler {
	return httphandler.NewStorageHandler(log, useCase)
}

type StorageUseCase interface {
	Upload(ctx context.Context, file entity.File) (*entity.UploadedInfo, error)
	Get(ctx context.Context, bucket string, filename string) (*entity.File, func(), error)
}

func NewStorageUseCase(log *zap.Logger, s3 S3) StorageUseCase {
	return usecase.NewStorageUseCase(log, s3)
}

type S3 interface {
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

func NewS3(
	log *zap.Logger,
	cfg *config.Config,
) (S3, error) {
	return repository.NewS3(
		log,
		cfg.S3.Endpoint,
		cfg.S3.AccessKeyID,
		cfg.S3.SecretAccessKey,
		cfg.S3.UseSSL,
		cfg.S3.Bucket,
	)
}

func NewConfig() (*config.Config, error) {
	return config.NewConfig()
}
