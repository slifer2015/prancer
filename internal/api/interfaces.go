package api

import "proj/internal/services/center"

type CenterService interface {
	Assign(in center.Point) error
	InitAgents()
}
