package bag

import "sync"

type ItemSet struct {
	Items    map[*Item]*Position
	occupied map[Position]bool
	mutex    sync.Mutex
}

// NewItemSet returns a new ItemSet.
func NewItemSet() *ItemSet {
	return &ItemSet{
		make(map[*Item]*Position, 0),
		make(map[Position]bool, 0),
		sync.Mutex{},
	}
}

// Add adds an item to this item set.
func (self *ItemSet) Add(item *Item) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	self.Items[item] = nil
}

// Remove removes an item from this item set.
func (self *ItemSet) Remove(item *Item) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	delete(self.Items, item)
}

// Count returns the number of items in this item set.
func (self *ItemSet) Count() int {
	return len(self.Items)
}

// Arrange arranges the items in this set using the given strategy.
func (self *ItemSet) Arrange(width, height int, strategy ArrangeStrategy) error {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	newItemSet, err := strategy.Arrange(width, height, self)
	if err != nil {
		return err
	}
	self.Items = newItemSet.Items
	self.occupied = newItemSet.occupied
	return nil
}

// Position sets the position of item to {x, y}.
func (self *ItemSet) Position(item *Item, x, y int) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	self.Items[item] = &Position{x, y}
	for i := 0; i < item.Height; i++ {
		for j := 0; j < item.Width; j++ {
			self.occupied[Position{x + j, y + i}] = true
		}
	}
}

// IsOccupied returns true if the position {x, y} is occupied.
func (self *ItemSet) IsOccupied(x, y int) bool {
	_, found := self.occupied[Position{x, y}]
	return found
}
