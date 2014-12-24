package bag

import "testing"

func TestSimpleStrategy(t *testing.T) {

	newTestBag := func(width, height int) *Bag {
		bag := NewBag(width, height)
		bag.Add(NewItem(1, 1, "Health Potion"))
		bag.Add(NewItem(1, 1, "Health Potion"))
		bag.Add(NewItem(1, 1, "Mana Potion"))
		bag.Add(NewItem(1, 1, "Mana Potion"))
		bag.Add(NewItem(1, 1, "Holy Hand Grenade of Antioch"))
		bag.Add(NewItem(1, 3, "Sword of a Thousand Truths"))
		bag.Add(NewItem(1, 3, "Excalibur"))
		bag.Add(NewItem(2, 3, "Mjölnir"))
		bag.Add(NewItem(3, 2, "BFG"))
		return bag
	}

	bag := newTestBag(20, 10)
	err := bag.Arrange(NewSimpleStrategy())
	if err != nil {
		t.Fatalf(`Expected nil, got "%v".`, err)
	}

	bag = newTestBag(5, 4)
	err = bag.Arrange(NewSimpleStrategy())
	if err == nil {
		t.Fatal(`Expected error, got nil`)
	}
}
