package bag

import (
	"fmt"
	"sort"
)

type SimpleStrategy struct {
}

// NewSimpleStrategy returns a new SimpleStrategy.
func NewSimpleStrategy() *SimpleStrategy {
	return &SimpleStrategy{}
}

// Arrange arranges the given itemSet as follows:
//    1. Sort the items by biggest to smallest
//    2. Iterate over each item:
//       2.1. Iterate over each empty slot in the space width x height until
//            one is found which can fit the current item.
//       2.2. If there isn't such a space, fail - we can't fit all the items
//            in.
// If d is max(width, height) and n is the number of items, this strategy
// runs in O(n * d^4).
func (self *SimpleStrategy) Arrange(width, height int, itemSet *ItemSet) (*ItemSet, error) {
	itemCount := itemSet.Count()
	itemList := make(ItemList, itemCount)
	i := 0
	for k, _ := range itemSet.items {
		itemList[i] = k
		i++
	}
	sort.Sort(itemList)
	newItemSet := NewItemSet()
	var canPlace bool
	for _, v := range itemList {
	PositionLoop:
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				// Can we place this item at this position?
				canPlace = true
			CanPlaceLoop:
				for ii := 0; ii < v.height; ii++ {
					for jj := 0; jj < v.width; jj++ {
						if newItemSet.IsOccupied(j+jj, i+ii) {
							canPlace = false
							break CanPlaceLoop
						}
					}
				}
				// If we can, great! Store the position and mark the slots
				// occupied by this item as occupied.
				if canPlace {
					newItemSet.Position(v, j, i)
					break PositionLoop
				}
			}
		}
		// We can't place this item, so the items in this bag don't fit using
		// this strategy.
		if !canPlace {
			return nil, fmt.Errorf(`Could not place item %v.`, v.payload)
		}
	}
	return newItemSet, nil
}
