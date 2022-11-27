<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# collection

```go
import "github.com/AdamHutchison/collection"
```

Go Collections are data structures that make filtering, mapping and general set iteration a breeze\!

## Index

- [type Collection](<#type-collection>)
  - [func (c *Collection[T]) All() []T](<#func-collectiont-all>)
  - [func (c *Collection[T]) Append(item T) Collection[T]](<#func-collectiont-append>)
  - [func (c *Collection[T]) Count() int](<#func-collectiont-count>)
  - [func (c *Collection[T]) Filter(fn func(item T) bool) *Collection[T]](<#func-collectiont-filter>)
  - [func (c *Collection[T]) First() T](<#func-collectiont-first>)
  - [func (c *Collection[T]) Get(index int) T](<#func-collectiont-get>)
  - [func (c *Collection[T]) GetNext() T](<#func-collectiont-getnext>)
  - [func (c *Collection[T]) Has(item T) bool](<#func-collectiont-has>)
  - [func (c *Collection[T]) HasNext() bool](<#func-collectiont-hasnext>)
  - [func (c *Collection[T]) Last() T](<#func-collectiont-last>)
  - [func (c *Collection[T]) Map(fn func(item T) T) *Collection[T]](<#func-collectiont-map>)
  - [func (c *Collection[T]) Merge(c2 Collection[T]) *Collection[T]](<#func-collectiont-merge>)
  - [func (c *Collection[T]) Pop() T](<#func-collectiont-pop>)
  - [func (c *Collection[T]) Shuffle() T](<#func-collectiont-shuffle>)


## type Collection

```go
type Collection[T any] struct {
    // contains filtered or unexported fields
}
```

### func \(\*Collection\[T\]\) All

```go
func (c *Collection[T]) All() []T
```

Returns the undelying collection slice.

### func \(\*Collection\[T\]\) Append

```go
func (c *Collection[T]) Append(item T) Collection[T]
```

Appends item to the end of the collection.

### func \(\*Collection\[T\]\) Count

```go
func (c *Collection[T]) Count() int
```

Returns the number of items currently in the collection.

### func \(\*Collection\[T\]\) Filter

```go
func (c *Collection[T]) Filter(fn func(item T) bool) *Collection[T]
```

Allows the filtering of a collection using a defined function. The function passed to Filter should accept a value of type T and return a bool, if the function returns true, the value will be kept in the collection, if the function returns false, the value will be removed.

### func \(\*Collection\[T\]\) First

```go
func (c *Collection[T]) First() T
```

Returns the first item in the collection

### func \(\*Collection\[T\]\) Get

```go
func (c *Collection[T]) Get(index int) T
```

Returns the value stored at the passed index

### func \(\*Collection\[T\]\) GetNext

```go
func (c *Collection[T]) GetNext() T
```

Used to retrieve the next item to in the iteration sequence. See \[collection.HasNext\] for details.

### func \(\*Collection\[T\]\) Has

```go
func (c *Collection[T]) Has(item T) bool
```

Determines if an element exists in the collection.

### func \(\*Collection\[T\]\) HasNext

```go
func (c *Collection[T]) HasNext() bool
```

Allows a collection to be iterated over as follows:

```
c := Collection[int]{
	set: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
}

converted := [string]{}

for c.HasNext() {
	got = append(converted, strconv.Itoa(c.GetNext()))
}
```

### func \(\*Collection\[T\]\) Last

```go
func (c *Collection[T]) Last() T
```

Returns the last item in the collection

### func \(\*Collection\[T\]\) Map

```go
func (c *Collection[T]) Map(fn func(item T) T) *Collection[T]
```

Allows a collection to be mapped into another using a function. The function provided to Map should accept an item of type T and return an item of type T.

### func \(\*Collection\[T\]\) Merge

```go
func (c *Collection[T]) Merge(c2 Collection[T]) *Collection[T]
```

Merges a given collection into the current. The second collection will be appended to the end of the current one.

### func \(\*Collection\[T\]\) Pop

```go
func (c *Collection[T]) Pop() T
```

Removes the last item in the collection and returns it.

### func \(\*Collection\[T\]\) Shuffle

```go
func (c *Collection[T]) Shuffle() T
```

Removes the first item in the collection and returns it. All othher item indexes are decreased by 1.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
