package _5_14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdder(t *testing.T) {
	a := assert.New(t)
	assertAdd := func(t testing.TB, got, want int) {
		t.Helper()
		a.Equal(want, got)
	}

	add := Adder(10)
	got := add(5)
	want := 15
	assertAdd(t, got, want)

	got = add(10)
	want = 25
	assertAdd(t, got, want)
}
