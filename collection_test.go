package collection

import (
	"reflect"
	"testing"
)

func TestCanGetAllItems(t *testing.T) {
	set := []int{1, 2, 3, 4}
	c := Collection[int]{set: set}

	want := set
	got := c.All()

	if reflect.DeepEqual(want, got) != true {
		t.Errorf("Wanted %v, Got %v", want, got)
	}
}

func TestCanAppendItemToCollection(t *testing.T) {
	c := Collection[string]{}
	c.Append("Test String")

	setLenth := len(c.All())

	if len(c.All()) != 1 {
		t.Errorf("Incorrect slice count when appending to collection. Wanted 1, Got %v", setLenth)
	}

	want := "Test String"
	got := c.All()[0]

	if want != got {
		t.Errorf("Wanted: %v, Got: %v", want, got)
	}
}

func TestCanPopItemFromCollection(t *testing.T) {
	c := Collection[int] {
		set: []int {1,2,3,4,5},
	}

	// Test the poped item is the last value in set
	want := 5
	got := c.Pop()

	if want != got {
		t.Errorf("Wanted: %v, Got: %v", want, got)
	}

	if len(c.All()) != 4 {
		t.Errorf("Failed to asserting collection had correct number of items")
	}

	// Test the item was removed from the collection
	valueWasRemovedFromCollection := reflect.DeepEqual([]int {1,2,3,4}, c.All())

	if !valueWasRemovedFromCollection {
		t.Errorf("Failed to assert that the value was removed from the collection")
	}
}
