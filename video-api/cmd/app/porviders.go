package main

import (
	"context"
	"fmt"
	"time"
	"video-api/config"
	"video-api/internal/server/amqp"
	"video-api/internal/usecase/transform"
	"video-api/internal/webapi/storage"
	"video-api/pkg/myresty"

	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewConfig() (*config.Config, error) {
	return config.NewConfig()
}

func NewQueue(lc fx.Lifecycle, log *zap.Logger, cfg *config.Config) (*rabbitmq.Connection, error) {

	conn, err := rabbitmq.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/",
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

			conn.Close()

			return nil
		},
	})

	return conn, nil

}

type Queue struct {
	ch   *rabbitmq.Channel
	msgs <-chan rabbitmq.Delivery
}

func NewVideoTransformQueue(
	lc fx.Lifecycle,
	log *zap.Logger,
	conn *rabbitmq.Connection,
	cfg *config.Config,
	handler *amqp.TransformHandler,
) (*Queue, error) {

	ch, err := conn.Channel()

	if err != nil {
		log.Sugar().Error(err.Error())
		return nil, err
	}

	q, err := ch.QueueDeclare(
		cfg.Transform.Queue, // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)

	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		return nil, err
	}

	go func() {
		for d := range msgs {
			log.Sugar().Infof("Received a message: %s", d.Body)
			handler.TransformVideo(d.Body)
		}
	}()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Sugar().Infof("amqp server start")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Sugar().Infof("connection  shutting")

			ch.Close()

			return nil
		},
	})

	return &Queue{
		ch:   ch,
		msgs: msgs,
	}, nil
}

func NewTransformHandler(log *zap.Logger, usecase *transform.VideoTransformer) *amqp.TransformHandler {
	return amqp.NewTransformHandler(log, usecase)
}

func NewVideoTransformUseCase(log *zap.Logger, storage *storage.StorageApi) *transform.VideoTransformer {
	return transform.New(log, storage)
}

func NewStorageApi(log *zap.Logger, cfg *config.Config) *storage.StorageApi {
	client := myresty.New(
		myresty.WithUrl(cfg.Storage.Url),
		myresty.WithTimeout(time.Duration(cfg.Storage.Timeout)*time.Millisecond),
	)

	return storage.New(log, client)
}
