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
		t.Errorf("Wanted: %v, Got: %v", want, got)
	}
}

func TestCanAppendItemToCollection(t *testing.T) {
	c := Collection[string]{}
	c.Append("Test String")

	setLenth := len(c.All())

	if len(c.All()) != 1 {
		t.Errorf("Incorrect slice count when appending to collection. Wanted: 1, Got: %v", setLenth)
	}

	want := "Test String"
	got := c.All()[0]

	if want != got {
		t.Errorf("Wanted: %v, Got: %v", want, got)
	}
}

func TestCanPopItemFromCollection(t *testing.T) {
	c := Collection[int]{
		set: []int{1, 2, 3, 4, 5},
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
	valueWasRemovedFromCollection := reflect.DeepEqual([]int{1, 2, 3, 4}, c.All())

	if !valueWasRemovedFromCollection {
		t.Errorf("Failed to assert that the value was removed from the collection")
	}
}

func TestCanSuffleCollection(t *testing.T) {
	c := Collection[int]{
		set: []int{1, 2, 3, 4, 5},
	}

	got := c.Shuffle()
	want := 1

	if want != got {
		t.Errorf("Shuffle() did not return the expected value. Wanted: %v, Got: %v", want, got)
	}

	newCollectionLength := len(c.All())

	if newCollectionLength != 4 {
		t.Errorf("Length of Collection after Shuffle() is not correct. Wanted: %v, Got: %v", 4, newCollectionLength)
	}

	valueWasRemovedFromCollection := reflect.DeepEqual([]int{2, 3, 4, 5}, c.All())

	if !valueWasRemovedFromCollection {
		t.Errorf("Failed asserting that 1st item was removed from Collection")
	}
}

func TestCanCountCollection(t *testing.T) {
	c := Collection[string]{
		set: []string{
			"item 1",
			"item 2",
			"item 3",
		},
	}

	want := 3
	got := c.Count()

	if want != got {
		t.Errorf("Wanted: %v, Got: %v", want, got)
	}
}

func TestCanFilterCollection(t *testing.T) {
	c := Collection[int]{
		set: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	want := []int{6, 7, 8, 9}

	got := c.Filter(func(item int) bool {
		return item > 5
	}).All()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Failed asserting that collection was filtered correctly. Wanted: %v, Got %v", want, got)
	}
}
