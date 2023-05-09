package probes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Readiness probe
// @Summary      Liveness probe
// @Description  Liveness probe
// @Tags         probe
// @Produce      json
// @Success      200
// @Router       /healthz/ready [get]
func NewReadinessProbe(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {

		var status int
		var content gin.H

		select {
		case <-ctx.Done():
			status = http.StatusGone
			content = gin.H{"status": "bad"}
		default:
			status = http.StatusOK
			content = gin.H{"status": "ok"}

		}

		c.IndentedJSON(status, content)
	}
}
