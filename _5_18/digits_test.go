package _5_18

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumOfDigits(t *testing.T) {
	a := assert.New(t)

	t.Run("", func(t *testing.T) {
		got := SumOfDigits(123)
		want := 6
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		got := SumOfDigits(100)
		want := 1
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		got := SumOfDigits(-123)
		want := 6
		a.Equal(want, got)
	})
}
