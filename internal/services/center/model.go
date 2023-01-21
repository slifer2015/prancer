package center

type Point struct {
	X float64
	Y float64
}

type AssignPointInput struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
