package collection

import (
	"reflect"
	"strconv"
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

func TestGetAllItemsReturnsCopy(t *testing.T) {
	set := []int{1, 2, 3, 4}
	c := Collection[int]{set: set}

	got := c.All()

	got[0] = 99

	if set[0] != 1 {
		t.Errorf("All() did not return a copy of the set")
	}
}

func TestCanAppendItemToCollection(t *testing.T) {
	c := Collection[string]{}
	c2 := c.Append("Test String")

	setLenth := len(c2.All())

	if len(c2.All()) != 1 {
		t.Errorf("Incorrect slice count when appending to collection. Wanted: 1, Got: %v", setLenth)
	}

	want := "Test String"
	got := c2.All()[0]

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
	c2, got := c.Pop()

	if want != got {
		t.Errorf("Wanted: %v, Got: %v", want, got)
	}

	if len(c2.All()) != 4 {
		t.Errorf("Failed to asserting collection had correct number of items")
	}

	// Test the item was removed from the collection
	valueWasRemovedFromCollection := reflect.DeepEqual([]int{1, 2, 3, 4}, c2.All())

	if !valueWasRemovedFromCollection {
		t.Errorf("Failed to assert that the value was removed from the collection")
	}
}

func TestCanSuffleCollection(t *testing.T) {
	c := Collection[int]{
		set: []int{1, 2, 3, 4, 5},
	}

	c2, got := c.Shuffle()
	want := 1

	if want != got {
		t.Errorf("Shuffle() did not return the expected value. Wanted: %v, Got: %v", want, got)
	}

	newCollectionLength := len(c2.All())

	if newCollectionLength != 4 {
		t.Errorf("Length of Collection after Shuffle() is not correct. Wanted: %v, Got: %v", 4, newCollectionLength)
	}

	valueWasRemovedFromCollection := reflect.DeepEqual([]int{2, 3, 4, 5}, c2.All())

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

func TestCanMapCollection(t *testing.T) {
	c := Collection[int]{
		set: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	want := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}

	got := c.Map(func(item int) int {
		return item + 1
	}).All()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Failed asserting that function was applied to each item in collection. Wanted: %v (%T), Got %v (%T)", want, want, got, got)
	}
}

func TestCollectionIsIterable(t *testing.T) {
	c := Collection[int]{
		set: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	want := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	got := []string{}

	for c.HasNext() {
		got = append(got, strconv.Itoa(c.GetNext()))
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Failed asserting that collection was iterated over correctly. Wanted: %v, Got %v", want, got)
	}
}

func TestResetsIterationWhenFinished(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Failed asserting that collection iteration was correctly reset. Error %v", r)
		}
	}()

	c := Collection[int]{
		set: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	for c.HasNext() {
		c.GetNext()
	}

	for c.HasNext() {
		c.GetNext()
	}
}

func TestHas(t *testing.T) {
	c := Collection[string]{
		set: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
	}

	want := true
	got := c.Has("4")

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}

	want = false
	got = c.Has("13")

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestCanGetValueAtIndex(t *testing.T) {
	c := Collection[string]{
		set: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
	}

	want := "5"
	got := c.Get(4)

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestCanGetFirstItemInCollection(t *testing.T) {
	c := Collection[int]{
		set: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	want := 1
	got := c.First()

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestCanGetLastItemInCollection(t *testing.T) {
	c := Collection[int]{
		set: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	want := 9
	got := c.Last()

	if want != got {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestCanMergeInSecondCollection(t *testing.T) {
	c := Collection[int]{
		set: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	c2 := Collection[int]{
		set: []int{10, 11, 12, 13, 14},
	}

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	got := c.Merge(c2)

	if reflect.DeepEqual(want, got) {
		t.Errorf("Failed asserting that second collection was correctly merged. Want: %v, Got: %v", want, got)
	}
}

func TestCanSliceCollection(t *testing.T) {
	c := Collection[int]{
		set: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	want := make([]int, 7)
	want = append(want, 3, 4, 5, 6, 7)
	got := c.Slice(2, 7).All()

	if reflect.DeepEqual(want, got) {
		t.Errorf("Failed asserting that second collection was correctly sliced. Want: %v (%T), Got: %v (%T)", want, want, got, got)
	}
}

func TestCanSortIntegersUsingAFunction(t *testing.T) {
	c := Collection[int]{
		set: []int{5, 2, 4, 3, 9, 1, 7, 8, 6},
	}

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	got := c.Sort(func(item1 int, item2 int) bool {
		return item1 < item2
	}).All()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestCanSortIntegersDescUsingAFunction(t *testing.T) {
	c := Collection[int]{
		set: []int{5, 2, 4, 3, 9, 1, 7, 8, 6},
	}

	want := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	got := c.Sort(func(item1 int, item2 int) bool {
		return item1 > item2
	}).All()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}

func TestSortHandlesEmptySlice(t *testing.T) {
	c := Collection[int]{}

	got := c.Sort(func(item1 int, item2 int) bool {
		return item1 > item2
	})

	if got.Count() != 0 {
		t.Errorf("Failed asserting that sorting an empty collection was handled correctly")
	}
}

func TestCanSortStringsUsingFunction(t *testing.T) {
	c := Collection[string]{
		set: []string{
			"AAAAAAA",
			"AAAAA",
			"A",
			"AA",
			"AAAA",
			"AAA",
			"AAAAAA",
		},
	}

	want := []string{
		"A",
		"AA",
		"AAA",
		"AAAA",
		"AAAAA",
		"AAAAAA",
		"AAAAAAA",
	}

	got := c.Sort(func(item1 string, item2 string) bool {
		return len(item1) < len(item2)
	}).All()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}
