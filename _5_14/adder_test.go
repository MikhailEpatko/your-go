package _5_14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdder(t *testing.T) {
	a := assert.New(t)
	add := Adder(10)

	t.Run("", func(t *testing.T) {
		got := add(5)
		want := 15
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		got := add(10)
		want := 25
		a.Equal(want, got)
	})
}
