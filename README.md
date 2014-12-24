bag
===

Sortable inventory for use in RPG-type games.

[![Build Status](https://travis-ci.org/NeilGarb/bag.svg?branch=master)](https://travis-ci.org/NeilGarb/bag)

Usage
=====

```
import (
    "fmt"
    "github.com/NeilGarb/bag"
)

s := bag.NewSimpleStrategy()
b := bag.NewBag(20, 10)
b.Add(NewItem(1, 1, "Holy Handgrenade of Antioch"))
b.Add(NewItem(1, 3, "Sword of a Thousand Truths"))
b.Add(NewItem(2, 2, "Mjölnir"))
b.Arrange(s)
for k, v := range b.ItemSet {
    fmt.Printf("Item %v is at {%d,%d}.", k.Payload, v.X, v.Y)
}
```

License
=======

You may use the code in any way you like, but you do so at your own risk.
