package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/millirud/easy-content/video-api/config"
	docs "github.com/millirud/easy-content/video-api/docs"
	"github.com/millirud/easy-content/video-api/internal/controller/http/http_metrics"
	"github.com/millirud/easy-content/video-api/internal/controller/http/middlewares"
	"github.com/millirud/easy-content/video-api/internal/controller/http/probes"
	"github.com/millirud/easy-content/video-api/pkg/httpserver"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/millirud/easy-content/video-api/internal/bootstrap"
	v1 "github.com/millirud/easy-content/video-api/internal/controller/http/v1"
	"golang.org/x/net/context"
)

func Run(cfg *config.Config) {
	var err error
	ctx, cancel := context.WithCancel(context.Background())

	handler := gin.New()
	handler.Use(middlewares.NewLogger(), gin.Recovery())

	handler.GET("/healthz", probes.NewLivenessProbe())
	handler.GET("/metrics", http_metrics.NewMetrics())

	setupSwagger(cfg.App)
	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	bs := bootstrap.New(cfg)

	v1.New(handler, bs)

	handler.GET("/healthz/ready", probes.NewReadinessProbe(ctx))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		cancel()
		fmt.Printf("app - Run - signal: %s /n", s.String())
	case err = <-httpServer.Notify():
		cancel()
		fmt.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		fmt.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}

func setupSwagger(cfg config.App) {
	docs.SwaggerInfo.Title = cfg.Name
	docs.SwaggerInfo.Description = cfg.Description
	docs.SwaggerInfo.Version = cfg.Version
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
