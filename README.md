bag
===

Sortable inventory for use in RPG-type games.

Use case:

```
s := NewBiggestToSmallestStrategy()
b := bag.NewBag(20, 10)
b.Add(NewItem(1, 1, "Holy Handgrenade of Antioch"))
b.Add(NewItem(1, 3, "Sword of a Thousand Truths"))
b.Add(NewItem(2, 2, "Mjölnir"))
b.Arrange(s)
// b.items is now a map from item -> position
```

_NB:_ This project is still under development.  You may use what code is here
in any way you choose, but you do so at your own risk.
