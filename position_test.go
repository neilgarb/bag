package bag

import "testing"

func TestPositionIsNeighbour(t *testing.T) {
	position := Position{10, 10}
	type test struct {
		position Position
		expected bool
	}
	tests := []test{
		{Position{0, 0}, false},
		{Position{10, 10}, false},
		{Position{9, 9}, false},
		{Position{10, 9}, true},
		{Position{11, 9}, false},
		{Position{9, 10}, true},
		{Position{10, 10}, false},
		{Position{11, 10}, true},
		{Position{9, 11}, false},
		{Position{10, 11}, true},
		{Position{11, 11}, false},
		{Position{20, 20}, false},
	}
	var result bool
	for _, test := range tests {
		result = position.IsNeighbour(test.position)
		if result != test.expected {
			t.Errorf(`Expected %v, got %v.`, test.expected, result)
		}
	}
}
