package bag

type ArrangeStrategy interface {
	Arrange(int, int, *ItemSet) (*ItemSet, error)
}
