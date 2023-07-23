package handler

import (
	"context"
	"net/http"
	"storage-api/internal/entity"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type (
	storageUseCase interface {
		Upload(ctx context.Context, file entity.File) (*entity.UploadedInfo, error)
	}
)

type storageHandler struct {
	log     *zap.Logger
	useCase storageUseCase
}

func NewStorageHandler(log *zap.Logger, useCase storageUseCase) *storageHandler {
	log.Info("NewStorageHandler")

	return &storageHandler{
		log:     log,
		useCase: useCase,
	}
}

func (s *storageHandler) Upload(ginCtx *gin.Context) {
	s.log.Info("storageHandler.upload")

	ctx := ginCtx.Request.Context()

	file, err := ginCtx.FormFile("file")

	if err != nil {
		ginCtx.AbortWithStatusJSON(http.StatusUnprocessableEntity, NewUnprocessableEntity(ctx, err.Error()))
		return
	}

	reader, err := file.Open()

	if err != nil {
		ginCtx.AbortWithStatusJSON(http.StatusUnprocessableEntity, NewUnprocessableEntity(ctx, err.Error()))
		return
	}

	defer reader.Close()

	uploaded, err := s.useCase.Upload(ctx, *entity.NewFile(
		reader,
		file.Size,
		file.Header.Get("Content-Type"),
		file.Filename,
	))

	if err != nil {
		ginCtx.AbortWithStatusJSON(http.StatusUnprocessableEntity, NewInternalServerError(ctx, err.Error()))
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"filename": uploaded.Filename,
		"bucket":   uploaded.Bucket,
	})

}

func (s *storageHandler) Get(ginCtx *gin.Context) {
	s.log.Info("storageHandler.Get")

	ginCtx.JSON(http.StatusOK, "get")

}
