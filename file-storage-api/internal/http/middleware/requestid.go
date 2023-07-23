package middleware

import (
	"net/http"
	"storage-api/pkg/requestid"

	httpHandler "storage-api/internal/http/handler"

	"github.com/gin-gonic/gin"
)

func NewRequestidMiddleware() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {

		reqId := ginCtx.Request.Header.Get("X-Request-Id")

		if reqId == "" {
			ginCtx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				httpHandler.NewUnprocessableEntity(ginCtx.Request.Context(), "request id header required"),
			)
			return
		}

		ginCtx.Request = ginCtx.Request.WithContext(
			requestid.WithRequestId(ginCtx.Request.Context(), reqId),
		)

		ginCtx.Next()
	}
}
