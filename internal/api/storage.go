package api

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Server struct {
	Router        *echo.Echo
	Logger        *zap.Logger
	CenterService CenterService
}
