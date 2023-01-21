package center

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"

	cmap "github.com/orcaman/concurrent-map"
	"go.uber.org/zap"
)

type service struct {
	logger *zap.Logger

	// points queue for income points
	points chan Point

	// lock main lock to lock main process (assigning agents)
	lock sync.RWMutex

	// agents store agents state using safe map
	agents cmap.ConcurrentMap
}

type sortByDistance []agent

func (s sortByDistance) Len() int {
	return len(s)
}
func (s sortByDistance) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortByDistance) Less(i, j int) bool {
	return s[i].distance < s[j].distance
}

func (s *service) Assign(p Point) {
	go func() {
		select {
		case s.points <- p:
		}
	}()
}

func (s *service) assign(p Point) {
	s.lock.Lock()
	var readyAgents = make([]agent, 0)
	for _, v := range s.agents.Items() {
		castedAgent, ok := v.(*agent)
		if !ok {
			panic("agent type not valid")
		}
		if castedAgent.ready {
			castedAgent.distance = s.calculateDistance(castedAgent.pos, p)
			readyAgents = append(readyAgents, *castedAgent)
		}
	}

	if len(readyAgents) == 0 {
		// we sleep for 1 second and requeue
		s.logger.Warn("no ready agents to take point")
		s.lock.Unlock()
		time.Sleep(1 * time.Second)
		s.points <- p
		return
	}

	sort.Sort(sortByDistance(readyAgents))

	assignedAgent := readyAgents[0]
	assignedAgent.ready = false
	s.agents.Set(assignedAgent.id, &assignedAgent)
	s.lock.Unlock()
	go func() {
		s.agentToPoint(assignedAgent.id, p)
	}()
}

func (s *service) Run() {
	s.logger.Info("main process started")
	for {
		select {
		case p := <-s.points:
			s.logger.Info(fmt.Sprintf("received point (%f,%f)", p.X, p.Y))
			s.assign(p)
		}
	}
}

func NewCenterService(logger *zap.Logger, startAgentsCount int) *service {
	agents := cmap.New()
	for i := 0; i < startAgentsCount; i++ {
		agents.Set(fmt.Sprint(i), &agent{
			id:    fmt.Sprint(i),
			ready: true,
			pos:   Point{0, 0},
		})
	}
	return &service{
		logger: logger,
		agents: agents,
		points: make(chan Point, 1000),
	}
}

func (s *service) calculateDistance(p1, p2 Point) float64 {
	tmp := math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2)
	return math.Pow(tmp, 0.5)
}
