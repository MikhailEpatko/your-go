package _9_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMagicSlice(t *testing.T) {
	a := assert.New(t)

	t.Run("", func(t *testing.T) {
		want := []int{24, 12, 8, 6}
		got := magicSlice([]int{1, 2, 3, 4})
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := []int{42, 35, 30, 210, 210}
		got := magicSlice([]int{5, 6, 7, 1, 1})
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := []int{0, -6, 0}
		got := magicSlice([]int{2, 0, -3})
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := []int{}
		got := magicSlice([]int{})
		a.Equal(want, got)
	})
}

func TestMagicString(t *testing.T) {
	a := assert.New(t)

	t.Run("", func(t *testing.T) {
		want := "[1, 2, 3, 4]"
		got := magicString([]int{1, 2, 3, 4})
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := "[]"
		got := magicString([]int{})
		a.Equal(want, got)
	})
}
