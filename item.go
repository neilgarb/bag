package bag

import "math"

type Item struct {
	Width  int
	Height int
	size   float64
	// payload can be used to link this item to an item in your application.
	Payload interface{}
}

// NewItem returns an new Item.
func NewItem(width, height int, payload interface{}) *Item {
	size := math.Sqrt(math.Pow(float64(width), 2) + math.Pow(float64(height), 2))
	return &Item{width, height, size, payload}
}
