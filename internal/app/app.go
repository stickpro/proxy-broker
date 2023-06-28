package app

import (
	"asocks-ws/internal/config"
	"asocks-ws/internal/repository"
	"asocks-ws/internal/router"
	"asocks-ws/internal/server"
	"asocks-ws/internal/service"
	"asocks-ws/pkg/database/mysql"
	"asocks-ws/pkg/logger"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configDir string) {
	cfg, err := config.Init(configDir)

	if err != nil {
		logger.Error(err)
		return
	}

	db, _ := mysql.ConnectionDataBase(cfg.DB.Host, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Port)

	repos := repository.NewRepositories(db)

	services := service.NewServices(service.Deps{
		Repository:  repos,
		KafkaConfig: cfg.Kafka,
	})

	router := router.NewRouter(services, cfg.ApiToken)
	srv := server.NewServer(cfg.HTTP, router.Init())

	if err != nil {
		logger.Error(err)
		return
	}

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	consumer := service.NewConsumer(cfg.Kafka, services)
	go func() {
		consumer.InitKafkaConsumer()
	}()

	logger.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}
