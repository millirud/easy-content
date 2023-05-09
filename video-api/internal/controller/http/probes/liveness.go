package probes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Liveness probe
// @Summary      Liveness probe
// @Description  Liveness probe
// @Tags         probe
// @Produce      json
// @Success      200
// @Router       /healthz [get]
func NewLivenessProbe() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}
