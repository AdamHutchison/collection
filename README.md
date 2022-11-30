<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# collection

```go
import "github.com/AdamHutchison/collection"
```

Go Collections are data structures that make filtering, mapping and general set iteration a breeze\!

## Index

- [type Collection](<#type-collection>)
  - [func (c *Collection[T]) All() []T](<#func-collectiont-all>)
  - [func (c *Collection[T]) Append(item T) *Collection[T]](<#func-collectiont-append>)
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
  - [func (c *Collection[T]) Pop() (*Collection[T], T)](<#func-collectiont-pop>)
  - [func (c *Collection[T]) Shuffle() (*Collection[T], T)](<#func-collectiont-shuffle>)
  - [func (c *Collection[T]) Slice(start int, limit int) *Collection[T]](<#func-collectiont-slice>)
  - [func (c *Collection[T]) Sort(fn func(item1 T, item2 T) bool) *Collection[T]](<#func-collectiont-sort>)


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

Returns the underlying collection slice.

### func \(\*Collection\[T\]\) Append

```go
func (c *Collection[T]) Append(item T) *Collection[T]
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

Used to retrieve the next item in the iteration sequence. See \[collection.HasNext\] for details.

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
func (c *Collection[T]) Pop() (*Collection[T], T)
```

Removes the last item in the collection and returns it.

### func \(\*Collection\[T\]\) Shuffle

```go
func (c *Collection[T]) Shuffle() (*Collection[T], T)
```

Removes the first item in the collection and returns it. All other item indexes are decreased by 1.

### func \(\*Collection\[T\]\) Slice

```go
func (c *Collection[T]) Slice(start int, limit int) *Collection[T]
```

Returns a collection containing a slice of the original values. The slice is inclusive of the start and exclusive of the limit.

### func \(\*Collection\[T\]\) Sort

```go
func (c *Collection[T]) Sort(fn func(item1 T, item2 T) bool) *Collection[T]
```

Allows a collection to be sorted using a defined function. The function passed to Sort is used to determine how the collection should be sorted The function accepts two items from the collection that can be compared to dictate the sorting order.

Sorting in ascending order: If you wish to sort the collection in ascending order then the function should return true if item1 should be placed lower than item2 e.g.

```
c := Collection[int]{
		set: []int{5, 2, 4, 3, 9, 1, 7, 8, 6},
	}

	sorted := c.Sort(func(item1 int, item2 int) bool {
		return item1 < item2
	}).All()
```

In the above example sorted would equal \[\]int\{1, 2, 3, 4, 5, 6, 7, 8, 9\}

Sorting in descending order: If you wish to sort the collection in descending order then the function should return true if item1 should be placed higher than item2 e.g.

```
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

sorted := c.Sort(func(item1 string, item2 string) bool {
	return len(item1) > len(item2)
}).All()
```

In the above scenario, sorted would equal:

```
[]string{
	"AAAAAAA",
	"AAAAAA",
	"AAAAA",
	"AAAA",
	"AAA",
	"AA",
	"A",
}
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
