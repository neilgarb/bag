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
	var canPlace, placed bool
	for _, v := range itemList {
		placed = false
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				// Can we place this item at this position?
				canPlace = true
				for ii := 0; ii < v.height; ii++ {
					for jj := 0; jj < v.width; jj++ {
						if newItemSet.IsOccupied(j+jj, i+ii) {
							canPlace = false
							break
						}
					}
					if !canPlace {
						break
					}
				}
				// If we can, great! Store the position and mark the slots
				// occupied by this item as occupied.
				if canPlace {
					newItemSet.items[v] = &Position{j, i}
					for ii := 0; ii < v.height; ii++ {
						for jj := 0; jj < v.width; jj++ {
							newItemSet.occupied[Position{j + jj, i + ii}] = true
						}
					}
					placed = true
					break
				}
			}
			if placed {
				break
			}
		}
		if !placed {
			return nil, fmt.Errorf(`Could not place item %v.`, v.payload)
		}
	}
	return newItemSet, nil
}
