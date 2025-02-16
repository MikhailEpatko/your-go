package _9_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumSlices(t *testing.T) {
	a := assert.New(t)

	t.Run("should return full slice", func(t *testing.T) {
		want := []int{2, 4, 6, 8}
		got := SumSlices([]int{1, 2, 3, 4}, []int{1, 2, 3, 4})
		a.Equal(want, got)
	})

	t.Run("when slices 2 is longer than slice 1 should return full slice", func(t *testing.T) {
		want := []int{2, 4, 6, 8}
		got := SumSlices([]int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5})
		a.Equal(want, got)
	})

	t.Run("when slices 1 is longer than slice 2 should return full slice", func(t *testing.T) {
		want := []int{2, 4, 6, 8}
		got := SumSlices([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4})
		a.Equal(want, got)
	})

	t.Run("when slices are empty should return empty slice", func(t *testing.T) {
		want := []int{}
		got := SumSlices([]int{}, []int{})
		a.Equal(want, got)
	})

	t.Run("when slices 2 is empty should return empty slice", func(t *testing.T) {
		want := []int{}
		got := SumSlices([]int{1, 2, 3, 4, 5}, []int{})
		a.Equal(want, got)
	})

	t.Run("when slice 1 is empty should return empty slice", func(t *testing.T) {
		want := []int{}
		got := SumSlices([]int{}, []int{1, 2, 3, 4, 5})
		a.Equal(want, got)
	})
}
