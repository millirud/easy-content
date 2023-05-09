package bootstrap

import (
	"os"

	"github.com/millirud/easy-content/video-api/config"
	"github.com/millirud/easy-content/video-api/pkg/logger"
	"github.com/rs/zerolog"
)

func New(cfg *config.Config) *Bootstrap {

	lg := logger.New(os.Stdin, cfg.Log.Level)

	return &Bootstrap{
		Logger: lg,
	}
}

type Bootstrap struct {
	Logger *zerolog.Logger
}
