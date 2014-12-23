package bag

import "math"

type Position struct {
	x int
	y int
}

// IsNeighbour returns true if p is adjacent to self in the x or y planes. A
// position is not its own neighbour.
func (self Position) IsNeighbour(p Position) bool {
	xdiff := math.Abs(float64(self.x) - float64(p.x))
	ydiff := math.Abs(float64(self.y) - float64(p.y))
	return xdiff+ydiff == 1
}
