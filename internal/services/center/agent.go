package center

import (
	"fmt"
	"time"

	cmap "github.com/orcaman/concurrent-map"
)

const (
	speedInMilliSecondsPerMeter float64 = 1000
)

type Agents struct {
	agents cmap.ConcurrentMap
}

type agent struct {
	id    string
	ready bool
	pos   Point

	// distance tmp fields only used to sort
	distance float64
}

func (s *service) agentToPoint(id string, p Point) {
	currentAgent, ok := s.agents.Get(id)
	if !ok {
		panic("agent not exists")
	}
	castedAgent, ok := currentAgent.(*agent)
	if !ok {
		panic("agent type not valid")
	}

	// make agent busy
	s.logger.Info(fmt.Sprintf("agents %s moving towards point\n , distance is = %f", castedAgent.id, castedAgent.distance))

	remainDistance := castedAgent.distance - float64(int(castedAgent.distance))
	if int(castedAgent.distance) >= 1 {
		times := int(castedAgent.distance)
		for i := 0; i < times; i++ {
			time.Sleep(time.Duration(speedInMilliSecondsPerMeter) * time.Millisecond)
			castedAgent.distance = castedAgent.distance - 1
			s.agents.Set(id, castedAgent)
			s.logger.Info(fmt.Sprintf("agents %s moving towards point , distance is = %f", castedAgent.id, castedAgent.distance))
		}
	}
	if remainDistance != 0 {
		time.Sleep(time.Duration(remainDistance*speedInMilliSecondsPerMeter) * time.Millisecond)
		castedAgent.distance = 0
		s.agents.Set(id, castedAgent)
		s.logger.Info(fmt.Sprintf("agents %s moving towards point , distance is = %f", castedAgent.id, castedAgent.distance))
	}

	castedAgent.pos = p
	castedAgent.ready = true
	s.agents.Set(id, castedAgent)
	s.logger.Info(fmt.Sprintf("agents %s is ready now", castedAgent.id))
}
