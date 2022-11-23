package collection

import (
	"reflect"
	"testing"
)

func TestCanGetSet(t *testing.T) {
	set := []int{1, 2, 3, 4}
	c := Collection[int]{set: set}

	want := set
	got := c.All()

	if reflect.DeepEqual(want, got) != true {
		t.Errorf("Wanted %v, Got %v", want, got)
	}
}
