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
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}
