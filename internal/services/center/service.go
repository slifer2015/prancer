package center

import (
	"math"
	"sort"
	"sync"
)

type service struct {
	agents Agents
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

func (s *service) Assign(p Point) error {
	var readyAgents = make([]agent, 0)
	for _, agent := range s.agents.agents {
		if !agent.ready {
			continue
		}
		distance := s.calculateDistance(agent.pos, p)
		agent.distance = distance
		readyAgents = append(readyAgents, agent)
	}
	sort.Sort(sortByDistance(readyAgents))
	// TODO : implement queue
	if len(readyAgents) == 0 {
		// agents not ready at the moment
	}

	assignedAgent := readyAgents[0]
	go func() {
		s.AgentToPoint(assignedAgent, p)
	}()
	return nil

}

func NewCenterService() *service {
	var startAgents = Agents{
		agents: make(map[int]agent, 0),
	}
	for i := 0; i < startAgentsCount; i++ {
		startAgents.agents[i] = agent{
			id:    i,
			ready: true,
			pos:   Point{0, 0},
		}
	}
	return &service{
		agents: Agents{
			agents: startAgents.agents,
			l:      sync.RWMutex{},
		},
	}
}

func (s *service) calculateDistance(p1, p2 Point) float64 {
	tmp := math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2)
	return math.Pow(tmp, 0.5)
}
