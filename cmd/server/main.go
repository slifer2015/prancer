package main

import (
	"github.com/labstack/echo/v4"

	"proj/internal/api"
	"proj/internal/services/center"
)

func main() {
	server := api.Server{
		Router:        echo.New(),
		CenterService: center.NewCenterService(),
	}

	server.CenterService.InitAgents()

	api.InitRoutes(server)
	server.Router.Start(":8080")
}
