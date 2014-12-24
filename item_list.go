package bag

type ItemList []*Item

func (self ItemList) Len() int {
	return len(self)
}

func (self ItemList) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self ItemList) Less(i, j int) bool {
	return self[i].size > self[j].size
}
