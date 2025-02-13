package _5_15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsWeekend(t *testing.T) {
	a := assert.New(t)
	asertResult := func(t testing.TB, got, want bool) {
		t.Helper()
		a.Equal(want, got)
	}

	got := IsWeekend(Monday)
	asertResult(t, got, false)

	got = IsWeekend(Tuesday)
	asertResult(t, got, false)

	got = IsWeekend(Wednesday)
	asertResult(t, got, false)

	got = IsWeekend(Thursday)
	asertResult(t, got, false)

	got = IsWeekend(Friday)
	asertResult(t, got, false)

	got = IsWeekend(Saturday)
	asertResult(t, got, true)

	got = IsWeekend(Sunday)
	asertResult(t, got, true)
}
