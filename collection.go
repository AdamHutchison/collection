// Go Collections are data structures that make filtering, mapping and general set iteration a breeze!
package collection

import "reflect"

type Collection[T any] struct {
	set   []T
	index int
}

// Returns the underlying collection slice.
func (c *Collection[T]) All() []T {
	return c.set
}

// Appends item to the end of the collection.
func (c *Collection[T]) Append(item T) Collection[T] {
	c.set = append(c.set, item)
	return *c
}

// Determines if an element exists in the collection.
func (c *Collection[T]) Has(item T) bool {
	for _, setItem := range c.set {
		if reflect.DeepEqual(setItem, item) {
			return true
		}
	}

	return false
}

// Returns the value stored at the passed index
func (c *Collection[T]) Get(index int) T {
	return c.set[index]
}

// Returns the first item in the collection
func (c *Collection[T]) First() T {
	return c.set[0]
}

// Returns the last item in the collection
func (c *Collection[T]) Last() T {
	return c.set[c.Count()-1]
}

// Removes the last item in the collection and returns it.
func (c *Collection[T]) Pop() T {
	item := c.set[len(c.set)-1]
	c.set = c.set[:len(c.set)-1]

	return item
}

// Removes the first item in the collection and returns it. All other item indexes are decreased by 1.
func (c *Collection[T]) Shuffle() T {
	item := c.set[0]
	c.set = c.set[1:len(c.set)]

	return item
}

// Returns the number of items currently in the collection.
func (c *Collection[T]) Count() int {
	return len(c.set)
}

// Allows the filtering of a collection using a defined function.
// The function passed to Filter should accept a value of type T and return a bool,
// if the function returns true, the value will be kept in the collection,
// if the function returns false, the value will be removed.
func (c *Collection[T]) Filter(fn func(item T) bool) *Collection[T] {
	newCollection := &Collection[T]{
		set: []T{},
	}

	for _, item := range c.set {
		if fn(item) {
			newCollection.Append(item)
		}
	}

	return newCollection
}

// Allows a collection to be mapped into another using a function.
// The function provided to Map should accept an item of type T and return an item of type T.
func (c *Collection[T]) Map(fn func(item T) T) *Collection[T] {
	newCollection := &Collection[T]{
		set: []T{},
	}

	for _, item := range c.set {
		newCollection.Append(fn(item))
	}

	return newCollection
}

// Allows a collection to be iterated over as follows:
//
//	c := Collection[int]{
//		set: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
//	}
//
//	converted := [string]{}
//
//	for c.HasNext() {
//		got = append(converted, strconv.Itoa(c.GetNext()))
//	}
func (c *Collection[T]) HasNext() bool {
	// Check if iteration is finished
	if c.index >= c.Count() {
		c.index = 0
		return false
	}

	return true
}

// Used to retrieve the next item in the iteration sequence.
// See [collection.HasNext] for details.
func (c *Collection[T]) GetNext() T {
	item := c.set[c.index]
	c.index++

	return item
}

// Merges a given collection into the current.
// The second collection will be appended to the end of the current one.
func (c *Collection[T]) Merge(c2 Collection[T]) *Collection[T] {
	c.set = append(c.set, c2.All()...)

	return c
}

// Returns a collection containing a slice of the original values.
// The slice is inclusive of the start and exclusive of the limit.
func (c *Collection[T]) Slice(start int, limit int) *Collection[T] {
	c.set = c.set[start:limit]

	return c
}

// Allows a collection to be sorted using a defined function.
// The function passed to Sort is used to determine how the collection should be sorted
// The function accepts two items from the collection that can be compared to dictate the sorting order.
//
// Sorting in ascending order:
// If you wish to sort the collection in ascending order then the function should
// return true if item1 should be placed lower than item2 e.g.
//
//	c := Collection[int]{
//			set: []int{5, 2, 4, 3, 9, 1, 7, 8, 6},
//		}
//
//		sorted := c.Sort(func(item1 int, item2 int) bool {
//			return item1 < item2
//		}).All()
//
// In the above example sorted would equal []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//
// Sorting in descending order:
// If you wish to sort the collection in descending order then the function should
// return true if item1 should be placed higher than item2 e.g.
//
//	c := Collection[string]{
//		set: []string{
//			"AAAAAAA",
//			"AAAAA",
//			"A",
//			"AA",
//			"AAAA",
//			"AAA",
//			"AAAAAA",
//		},
//	 }
//
//	sorted := c.Sort(func(item1 string, item2 string) bool {
//		return len(item1) > len(item2)
//	}).All()
//
// In the above scenario, sorted would equal:
//
//	[]string{
//		"AAAAAAA",
//		"AAAAAA",
//		"AAAAA",
//		"AAAA",
//		"AAA",
//		"AA",
//		"A",
//	}
func (c *Collection[T]) Sort(fn func(item1 T, item2 T) bool) *Collection[T] {
	c.set = sortSlice(c.set, fn)

	return c
}

func sortSlice[T any](items []T, fn func(item1 T, item2 T) bool) []T {
	// Check if sorting is required
	if len(items) < 2 {
		return items
	}

	pivot := items[0]

	lowerSlice, upperSlice := partitionSlice(pivot, items[1:], fn)

	sortedItems := append(sortSlice(lowerSlice, fn), pivot)

	return append(sortedItems, sortSlice(upperSlice, fn)...)
}

func partitionSlice[T any](pivot T, items []T, fn func(item1 T, item2 T) bool) (lowerSlice []T, upperSlice []T) {
	for _, item := range items {
		if fn(pivot, item) {
			upperSlice = append(upperSlice, item)
		} else {
			lowerSlice = append(lowerSlice, item)
		}
	}

	return
}
