package _8_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindromeStr(t *testing.T) {
	a := assert.New(t)
	assertStr := func(t testing.TB, want, got string) {
		t.Helper()
		a.Equal(want, got)
	}

	want := SUCCESS_MSG
	got := isPalindromeStr([10]int{1, 2, 3, 2, 1, 1, 2, 3, 2, 1})
	assertStr(t, want, got)

	want = FAIL_MSG
	got = isPalindromeStr([10]int{1, 2, 3, 4, 5, 4, 3, 2, 1, 1})
	assertStr(t, want, got)
}
