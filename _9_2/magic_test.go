package _9_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMagicSlice(t *testing.T) {
	a := assert.New(t)
	assertSlices := func(t *testing.T, want, got []int) {
		t.Helper()
		a.Equal(want, got)
	}

	want := []int{24, 12, 8, 6}
	got := magicSlice([]int{1, 2, 3, 4})
	assertSlices(t, want, got)

	want = []int{42, 35, 30, 210, 210}
	got = magicSlice([]int{5, 6, 7, 1, 1})
	assertSlices(t, want, got)

	want = []int{0, -6, 0}
	got = magicSlice([]int{2, 0, -3})
	assertSlices(t, want, got)

	want = []int{}
	got = magicSlice([]int{})
	assertSlices(t, want, got)
}

func TestMagicString(t *testing.T) {
	a := assert.New(t)
	assertStrings := func(t *testing.T, want, got string) {
		t.Helper()
		a.Equal(want, got)
	}

	want := "[1, 2, 3, 4]"
	got := magicString([]int{1, 2, 3, 4})
	assertStrings(t, want, got)

	want = "[]"
	got = magicString([]int{})
	assertStrings(t, want, got)
}
