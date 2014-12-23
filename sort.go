package bag

type SortStrategy interface {
	Arrange(int, int, map[*Item]*Position) map[*Item]*Position
}
