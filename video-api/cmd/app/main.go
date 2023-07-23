package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			zap.NewProduction,
			NewConfig,
			NewQueue,
		),
		fx.Invoke(func(*amqp.Connection) {}),
	).Run()
}
