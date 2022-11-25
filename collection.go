package collection

type Collection[T any] struct {
	set []T
}

func (c *Collection[T]) All() []T {
	return c.set
}

func (c *Collection[T]) Append(item T) Collection[T] {
	c.set = append(c.set, item)
	return *c
}

func (c *Collection[T]) Pop() T {
	item := c.set[len(c.set) - 1]
	c.set = c.set[:len(c.set) - 1]

	return item
}
