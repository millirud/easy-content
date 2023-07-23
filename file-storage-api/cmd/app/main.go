package main

import (
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			zap.NewProduction,
			NewHTTPServer,
			NewGin,
			NewStorageHandler,
			NewStorageUseCase,
			NewS3,
			NewConfig,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
