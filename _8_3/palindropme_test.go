package _8_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindromeStr(t *testing.T) {
	a := assert.New(t)

	t.Run("", func(t *testing.T) {
		want := SUCCESS_MSG
		got := isPalindromeStr([10]int{1, 2, 3, 2, 1, 1, 2, 3, 2, 1})
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := FAIL_MSG
		got := isPalindromeStr([10]int{1, 2, 3, 4, 5, 4, 3, 2, 1, 1})
		a.Equal(want, got)
	})
}
