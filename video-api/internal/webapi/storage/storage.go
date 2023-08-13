package storage

import (
	"errors"
	"io"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type StorageApi struct {
	log    *zap.Logger
	client *resty.Client
}

func New(
	log *zap.Logger,
	client *resty.Client,
) *StorageApi {
	return &StorageApi{
		log:    log,
		client: client,
	}
}

var (
	ErrorFailedRequest = errors.New("failed request")
)

func (s *StorageApi) Download(bucket string, filename string) (io.Reader, func() error, error) {

	resp, err := s.client.R().
		SetPathParams(map[string]string{
			"bucket":   bucket,
			"filename": filename,
		}).
		SetHeader("X-Request-Id", "fake-request-id"). //TODO
		Get("/v1/view/{bucket}/{filename}")

	if err != nil {
		s.log.Sugar().Error(err.Error())
		return nil, nil, err
	}

	if resp.IsError() {
		s.log.Sugar().Errorf("request %s response: %d", resp.Request.URL, resp.StatusCode())
		return nil, nil, ErrorFailedRequest
	}

	return resp.RawResponse.Body, func() error { return resp.RawResponse.Body.Close() }, nil
}
