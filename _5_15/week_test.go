package _5_15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsWeekend(t *testing.T) {
	a := assert.New(t)

	t.Run("", func(t *testing.T) {
		got := IsWeekend(Monday)
		a.Equal(false, got)
	})

	t.Run("", func(t *testing.T) {
		got := IsWeekend(Tuesday)
		a.Equal(false, got)
	})

	t.Run("", func(t *testing.T) {
		got := IsWeekend(Wednesday)
		a.Equal(false, got)
	})

	t.Run("", func(t *testing.T) {
		got := IsWeekend(Thursday)
		a.Equal(false, got)
	})

	t.Run("", func(t *testing.T) {
		got := IsWeekend(Friday)
		a.Equal(false, got)
	})

	t.Run("", func(t *testing.T) {
		got := IsWeekend(Saturday)
		a.Equal(true, got)
	})

	t.Run("", func(t *testing.T) {
		got := IsWeekend(Sunday)
		a.Equal(true, got)
	})
}
