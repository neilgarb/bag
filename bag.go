package bag

import (
	"bytes"
	"sync"
)

type Bag struct {
	width   int
	height  int
	itemSet *ItemSet
	mutex   sync.Mutex
}

// NewBag returns a new Bag.
func NewBag(width, height int) *Bag {
	return &Bag{width, height, NewItemSet(), sync.Mutex{}}
}

// String returns a string representation of this bag as a 2D matrix of 0s and
// 1s, where a 0 is an empty slot and a 1 is an occupied slot.
func (self *Bag) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("\n")
	for i := 0; i < self.height; i++ {
		for j := 0; j < self.width; j++ {
			if self.itemSet.IsOccupied(j, i) {
				buf.WriteString("1")
			} else {
				buf.WriteString("0")
			}
			if j != self.width-1 {
				buf.WriteString(" ")
			}
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

// Add adds an item to this bag. Note that adding an item a second time
// will have no effect. The item is added with nil position, which means it
// hasn't been positioned yet.
func (self *Bag) Add(item *Item) {
	self.itemSet.Add(item)
}

// Remove removes the given item from the bag.
func (self *Bag) Remove(item *Item) {
	self.itemSet.Remove(item)
}

// Count returns the number of items in this bag.
func (self *Bag) Count() int {
	return self.itemSet.Count()
}

// Arrange will give each item in the bag a position using the given strategy.
func (self *Bag) Arrange(strategy ArrangeStrategy) error {
	return self.itemSet.Arrange(self.width, self.height, strategy)
}
