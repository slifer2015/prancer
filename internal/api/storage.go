package api

import "github.com/labstack/echo/v4"

type Server struct {
	Router *echo.Echo

	CenterService CenterService
}
