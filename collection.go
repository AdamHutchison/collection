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
