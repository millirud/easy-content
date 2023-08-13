package amqp

import (
	"encoding/json"
	"video-api/internal/entity"

	"go.uber.org/zap"
)

type TransformHandler struct {
	log         *zap.Logger
	transformer videoTransformer
}

func NewTransformHandler(
	log *zap.Logger,
	transformer videoTransformer,
) *TransformHandler {
	return &TransformHandler{
		log:         log,
		transformer: transformer,
	}
}

type (
	videoTransformer interface {
		TransformVideo(fileInfo entity.UploadedFile) error
	}
)

func (t *TransformHandler) TransformVideo(payload []byte) {
	t.log.Sugar().Infof("TransformHandler.TransformVideo")

	parsedPayload := videoTransformPayload{}
	json.Unmarshal(payload, &parsedPayload)

	t.transformer.TransformVideo(entity.NewUploadedFile(
		parsedPayload.Bucket,
		parsedPayload.Filename,
	))
}
