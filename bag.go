package bag

import (
	"bytes"
	"sync"
)

type Bag struct {
	Width   int
	Height  int
	ItemSet *ItemSet
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
	for i := 0; i < self.Height; i++ {
		for j := 0; j < self.Width; j++ {
			if self.ItemSet.IsOccupied(j, i) {
				buf.WriteString("1")
			} else {
				buf.WriteString("0")
			}
			if j != self.Width-1 {
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
	self.ItemSet.Add(item)
}

// Remove removes the given item from the bag.
func (self *Bag) Remove(item *Item) {
	self.ItemSet.Remove(item)
}

// Count returns the number of items in this bag.
func (self *Bag) Count() int {
	return self.ItemSet.Count()
}

// Arrange will give each item in the bag a position using the given strategy.
func (self *Bag) Arrange(strategy ArrangeStrategy) error {
	return self.ItemSet.Arrange(self.Width, self.Height, strategy)
}
