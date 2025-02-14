package _7_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlay(t *testing.T) {
	a := assert.New(t)
	assertNum := func(t testing.TB, want, got, cnt int) {
		t.Helper()
		a.Equal(want, got)
		a.True(cnt <= 6)
	}

	for i := 1; i <= 100; i++ {
		cnt = 0
		val = i
		got := play()
		assertNum(t, val, got, cnt)
	}
}
