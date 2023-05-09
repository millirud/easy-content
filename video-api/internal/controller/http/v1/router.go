package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/millirud/easy-content/video-api/internal/bootstrap"
	"github.com/rs/zerolog"
)

func New(handler *gin.Engine, bs *bootstrap.Bootstrap) {
	r := Router{
		logger: bs.Logger,
	}

	handler.POST("/v1/convert", r.conver)
}

type Router struct {
	logger *zerolog.Logger
}

func (r *Router) conver(c *gin.Context) {

	c.IndentedJSON(http.StatusNotImplemented, "wait")
}
