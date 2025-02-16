package _5_22

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMovePirateDone(t *testing.T) {
	a := assert.New(t)
	stepCnt = 0

	t.Run("", func(t *testing.T) {
		want := "Пират переместился на плиту 1"
		got, _ := movePirateMsg(false)
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := "Пират переместился на плиту 2"
		got, _ := movePirateMsg(false)
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := "Пират переместился на плиту 3\nПират преодолел все ловушки"
		got, _ := movePirateMsg(false)
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := "game over"
		_, err := movePirateMsg(false)
		a.NotNil(err)
		a.Equal(want, err.Error())
	})
}

func TestMovePirateKilled(t *testing.T) {
	a := assert.New(t)
	stepCnt = 0

	t.Run("", func(t *testing.T) {
		want := "Пират переместился на плиту 1\nПират ранен"
		got, _ := movePirateMsg(true)
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := "Пират переместился на плиту 2\nПират убит"
		got, _ := movePirateMsg(true)
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := "game over"
		_, err := movePirateMsg(true)
		a.NotNil(err)
		a.Equal(want, err.Error())
	})
}
