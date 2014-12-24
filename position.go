package bag

import "math"

type Position struct {
	X int
	Y int
}

// IsNeighbour returns true if p is adjacent to self in the x or y planes. A
// position is not its own neighbour.
func (self Position) IsNeighbour(p Position) bool {
	xdiff := math.Abs(float64(self.X) - float64(p.X))
	ydiff := math.Abs(float64(self.Y) - float64(p.Y))
	return xdiff+ydiff == 1
}
