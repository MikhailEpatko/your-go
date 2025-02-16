package _7_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuffix1(t *testing.T) {
	a := assert.New(t)

	t.Run("", func(t *testing.T) {
		want := "ся"
		got := suffix1(1)
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := "ся"
		got := suffix1(21)
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := "ся"
		got := suffix1(131)
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := "ось"
		got := suffix1(2)
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := "ось"
		got := suffix1(11)
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := "ось"
		got := suffix1(23)
		a.Equal(want, got)
	})
}

func TestSuffix2(t *testing.T) {
	a := assert.New(t)

	t.Run("", func(t *testing.T) {
		want := "ок"
		got := suffix2(1)
		a.Equal(want, got)
	})

	for i := 2; i <= 4; i++ {
		t.Run("", func(t *testing.T) {
			want := "ка"
			got := suffix2(i)
			a.Equal(want, got)
		})
	}

	for i := 5; i <= 20; i++ {
		t.Run("", func(t *testing.T) {
			want := "ков"
			got := suffix2(i)
			a.Equal(want, got)
		})
	}

	t.Run("", func(t *testing.T) {
		want := "ок"
		got := suffix2(21)
		a.Equal(want, got)
	})

	for i := 22; i <= 24; i++ {
		t.Run("", func(t *testing.T) {
			want := "ка"
			got := suffix2(i)
			a.Equal(want, got)
		})
	}

	for i := 25; i <= 30; i++ {
		t.Run("", func(t *testing.T) {
			want := "ков"
			got := suffix2(i)
			a.Equal(want, got)
		})
	}
}

func TestRollDiseMsg(t *testing.T) {
	a := assert.New(t)

	t.Run("", func(t *testing.T) {
		want := "Выпало 4 и 4, в сумме 8, на это потребовался 1 бросок."
		got := rollDiceMsg(4, 4, 8, 8, 1)
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := "Выпало 1 и 6, в сумме 7, бросаем еще раз."
		got := rollDiceMsg(1, 6, 7, 8, 1)
		a.Equal(want, got)
	})

	t.Run("", func(t *testing.T) {
		want := "Выпало 5 и 3, в сумме 8, на это потребовалось 5 бросков."
		got := rollDiceMsg(5, 3, 8, 8, 5)
		a.Equal(want, got)
	})
}

func TestRollDise(t *testing.T) {
	rollDice(8)
}
