package _5_22

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMovePirateDone(t *testing.T) {
	a := assert.New(t)
	assertMoveMsg := func(t testing.TB, want, got string) {
		t.Helper()
		a.Equal(want, got)
	}
	stepCnt = 0

	want := "Пират переместился на плиту 1"
	got, _ := movePirateMsg(false)
	assertMoveMsg(t, want, got)

	want = "Пират переместился на плиту 2"
	got, _ = movePirateMsg(false)
	assertMoveMsg(t, want, got)

	want = "Пират переместился на плиту 3\nПират преодолел все ловушки"
	got, _ = movePirateMsg(false)
	assertMoveMsg(t, want, got)

	want = "game over"
	_, err := movePirateMsg(false)
	a.NotNil(err)
	assertMoveMsg(t, want, err.Error())
}

func TestMovePirateKilled(t *testing.T) {
	a := assert.New(t)
	assertMoveMsg := func(t testing.TB, want, got string) {
		t.Helper()
		a.Equal(want, got)
	}
	stepCnt = 0

	want := "Пират переместился на плиту 1\nПират ранен"
	got, _ := movePirateMsg(true)
	assertMoveMsg(t, want, got)

	want = "Пират переместился на плиту 2\nПират убит"
	got, _ = movePirateMsg(true)
	assertMoveMsg(t, want, got)

	want = "game over"
	_, err := movePirateMsg(true)
	a.NotNil(err)
	assertMoveMsg(t, want, err.Error())
}
