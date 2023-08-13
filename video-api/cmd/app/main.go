package main

import (
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			zap.NewProduction,
			NewConfig,
			NewQueue,
			NewVideoTransformQueue,
			NewTransformHandler,
			NewVideoTransformUseCase,
			NewStorageApi,
		),
		fx.Invoke(
			func(*rabbitmq.Connection) {},
			func(*Queue) {},
		),
	).Run()
}
