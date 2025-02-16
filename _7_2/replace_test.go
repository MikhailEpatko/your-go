package _7_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplceStr(t *testing.T) {
	a := assert.New(t)

	t.Run("", func(t *testing.T) {
		want := "Какашка!"
		got := replacedStr("Кукушка!")
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := "Лажа."
		got := replacedStr("Лужа.")
		a.Equal(want, got)
	})
}
