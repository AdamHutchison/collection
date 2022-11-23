package collection

type Collection[T any] struct {
	set []T
}

func (c *Collection[T]) All() []T {
	return c.set
}
