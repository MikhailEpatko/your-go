package _8_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumNeighbors(t *testing.T) {
	a := assert.New(t)

	t.Run("", func(t *testing.T) {
		want := [ARRAAY_LENGTH]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 9}
		got := SumNeighbors([ARRAAY_LENGTH]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		a.Equal(want, got)
	})
}
