package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"proj/internal/services/center"
)

// handleAssignPoint responsible to assign new point to agents
func handleAssignPoint(centerService CenterService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var in center.AssignPointInput
		err := c.Bind(&in)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		err = centerService.Assign(center.Point{
			X: in.X,
			Y: in.Y,
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, nil)
	}
}

func InitRoutes(s Server) {
	s.Router.POST("/assign", handleAssignPoint(s.CenterService))
}
