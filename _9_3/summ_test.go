package _9_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumSlices(t *testing.T) {
	a := assert.New(t)
	assertSlices := func(t *testing.T, want, got []int) {
		t.Helper()
		a.Equal(want, got)
	}

	want := []int{2, 4, 6, 8}
	got := SumSlices([]int{1, 2, 3, 4}, []int{1, 2, 3, 4})
	assertSlices(t, want, got)

	want = []int{2, 4, 6, 8}
	got = SumSlices([]int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5})
	assertSlices(t, want, got)

	want = []int{2, 4, 6, 8}
	got = SumSlices([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4})
	assertSlices(t, want, got)

	want = []int{}
	got = SumSlices([]int{}, []int{})
	assertSlices(t, want, got)

	want = []int{}
	got = SumSlices([]int{1, 2, 3, 4, 5}, []int{})
	assertSlices(t, want, got)

	want = []int{}
	got = SumSlices([]int{}, []int{1, 2, 3, 4, 5})
	assertSlices(t, want, got)
}
