package bag

import "testing"

func TestBagString(t *testing.T) {
	bag := NewBag(1, 1)
	expected := "\n0\n"
	result := bag.String()
	if result != expected {
		t.Errorf(`Expected "%v", got "%v".`, expected, result)
	}

	bag = NewBag(2, 2)
	expected = "\n0 0\n0 0\n"
	result = bag.String()
	if result != expected {
		t.Errorf(`Expected "%v", got "%v".`, expected, result)
	}

	bag.ItemSet.occupied[Position{1, 1}] = true // NOT the use case, but useful
	expected = "\n0 0\n0 1\n"
	result = bag.String()
	if result != expected {
		t.Errorf(`Expected "%v", got "%v".`, expected, result)
	}
}

func TestBagAddRemoveCount(t *testing.T) {
	bag := NewBag(1, 1)
	expected := 0
	result := bag.Count()
	if result != expected {
		t.Errorf(`Expected %v, got %v.`, expected, result)
	}

	item := NewItem(1, 1, nil)
	bag.Add(item)
	expected = 1
	result = bag.Count()
	if result != expected {
		t.Errorf(`Expected %v, got %v.`, expected, result)
	}

	bag.Add(item)
	expected = 1
	result = bag.Count()
	if result != expected {
		t.Errorf(`Expected %v, got %v.`, expected, result)
	}

	bag.Remove(item)
	expected = 0
	result = bag.Count()
	if result != expected {
		t.Errorf(`Expected %v, got %v.`, expected, result)
	}
}
