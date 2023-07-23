package main

import (
	"context"
	"fmt"
	"video-api/config"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewConfig() (*config.Config, error) {
	return config.NewConfig()
}

func NewQueue(lc fx.Lifecycle, log *zap.Logger, cfg *config.Config) (*amqp.Connection, error) {

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/",
		cfg.RabbitMq.User,
		cfg.RabbitMq.Password,
		cfg.RabbitMq.Host,
		cfg.RabbitMq.Port,
	))

	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("start queue")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("shutting queue")

			return nil
		},
	})

	return conn, nil

}
