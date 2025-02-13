package _5_18

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumOfDigits(t *testing.T) {
	a := assert.New(t)

	got := SumOfDigits(123)
	want := 6
	a.Equal(want, got)

	got = SumOfDigits(100)
	want = 1
	a.Equal(want, got)

	got = SumOfDigits(-123)
	want = 6
	a.Equal(want, got)
}
