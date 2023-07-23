package handler

import (
	"context"
	"net/http"
	"storage-api/pkg/requestid"
)

type HttpError struct {
	RequestId   string `json:"requestId"`
	Status      string `json:"status"`
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func NewUnprocessableEntity(ctx context.Context, description string) *HttpError {
	return &HttpError{
		RequestId:   requestid.GetRequestId(ctx),
		Code:        http.StatusUnprocessableEntity,
		Status:      "UnprocessableEntity",
		Description: description,
	}
}

func NewHttpError(requestId string, status string, code int, description string) *HttpError {
	return &HttpError{
		RequestId:   requestId,
		Status:      status,
		Code:        code,
		Description: description,
	}
}

func NewInternalServerError(ctx context.Context, description string) *HttpError {
	return &HttpError{
		RequestId:   requestid.GetRequestId(ctx),
		Code:        http.StatusInternalServerError,
		Status:      "UnprocessableEntity",
		Description: description,
	}
}
