package main

import (
	golog "log"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"proj/internal/api"
	"proj/internal/config"
	"proj/internal/services/center"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		golog.Fatal(err)
	}
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	server := api.Server{
		Logger:        logger,
		Router:        echo.New(),
		CenterService: center.NewCenterService(logger, cfg.StartAgentsCount),
	}
	api.InitRoutes(server)
	logger.Info("server started", zap.String("port", cfg.Port))

	// main process
	go func() {
		server.CenterService.Run()
	}()

	server.Router.Start(cfg.Port)
}
