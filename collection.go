// Go Collections are data structures that make filtering, mapping and general set iteration a breeze!
package collection

type Collection[T any] struct {
	set []T
}

// Returns the undelying collection slice
func (c *Collection[T]) All() []T {
	return c.set
}

// Appends item to the end of the collection
func (c *Collection[T]) Append(item T) Collection[T] {
	c.set = append(c.set, item)
	return *c
}

// Removes the last item in the collection and returns it.
func (c *Collection[T]) Pop() T {
	item := c.set[len(c.set)-1]
	c.set = c.set[:len(c.set)-1]

	return item
}

// Removes the first item in the collection and returns it. All othher item indexes are decreased by 1.
func (c *Collection[T]) Shuffle() T {
	item := c.set[0]
	c.set = c.set[1:len(c.set)]

	return item
}

// Returns the number of items currently in the collection
func (c *Collection[T]) Count() int {
	return len(c.set)
}
