package center

import (
	"fmt"
	"sync"
	"time"
)

const (
	startAgentsCount = 10
)

type Agents struct {
	agents map[int]agent
	l      sync.RWMutex
}

type agent struct {
	id    int
	ready bool
	pos   Point

	// distance tmp fields only used to sort
	distance float64
}

func (a *Agents) SetReadiness(id int, ready bool) {
	a.l.Lock()
	defer a.l.Unlock()
	tmp := a.agents[id]
	tmp.ready = ready
	a.agents[id] = tmp
}

func (a *Agents) SetPoint(id int, p Point) {
	a.l.Lock()
	defer a.l.Unlock()
	tmp := a.agents[id]
	tmp.pos = p
	a.agents[id] = tmp
}

func (s *service) AgentToPoint(a agent, p Point) {
	s.agents.SetReadiness(a.id, false)
	fmt.Printf("agents %d moving towards point\n", a.id)
	// TODO : for half second
	for i := 0; i < int(a.distance); i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("agents %d moving to point, distance : %d\n", a.id, int(a.distance)-(i+1))
	}

	s.agents.SetPoint(a.id, p)
	s.agents.SetReadiness(a.id, true)
	fmt.Printf("agents %d is ready now\n", a.id)
}

// InitAgents initialize agents at point (0,0)
func (s *service) InitAgents() {

}
