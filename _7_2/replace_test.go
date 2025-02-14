package _7_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplceStr(t *testing.T) {
	a := assert.New(t)
	assertString := func(t testing.TB, want, got string) {
		t.Helper()
		a.Equal(want, got)
	}

	want := "Какашка!"
	got := replacedStr("Кукушка!")
	assertString(t, want, got)

	want = "Лажа."
	got = replacedStr("Лужа.")
	assertString(t, want, got)
}
