package _7_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuffix1(t *testing.T) {
	a := assert.New(t)
	assertMsg := func(t testing.TB, want, got string) {
		t.Helper()
		a.Equal(want, got)
	}

	want := "ся"
	got := suffix1(1)
	assertMsg(t, got, want)

	want = "ся"
	got = suffix1(21)
	assertMsg(t, got, want)

	want = "ся"
	got = suffix1(131)
	assertMsg(t, got, want)

	want = "ось"
	got = suffix1(2)
	assertMsg(t, got, want)

	want = "ось"
	got = suffix1(11)
	assertMsg(t, got, want)

	want = "ось"
	got = suffix1(23)
	assertMsg(t, got, want)
}

func TestSuffix2(t *testing.T) {
	a := assert.New(t)
	assertMsg := func(t testing.TB, want, got string) {
		t.Helper()
		a.Equal(want, got)
	}

	want := "ок"
	got := suffix2(1)
	assertMsg(t, got, want)

	for i := 2; i <= 4; i++ {
		want = "ка"
		got = suffix2(i)
		assertMsg(t, got, want)
	}

	for i := 5; i <= 20; i++ {
		want = "ков"
		got = suffix2(i)
		assertMsg(t, got, want)
	}

	want = "ок"
	got = suffix2(21)
	assertMsg(t, got, want)

	for i := 22; i <= 24; i++ {
		want = "ка"
		got = suffix2(i)
		assertMsg(t, got, want)
	}

	for i := 25; i <= 30; i++ {
		want = "ков"
		got = suffix2(i)
		assertMsg(t, got, want)
	}
}

func TestRollDiseMsg(t *testing.T) {
	a := assert.New(t)
	assertMsg := func(t testing.TB, want, got string) {
		t.Helper()
		a.Equal(want, got)
	}

	want := "Выпало 4 и 4, в сумме 8, на это потребовался 1 бросок."
	got := rollDiceMsg(4, 4, 8, 8, 1)
	assertMsg(t, got, want)

	want = "Выпало 1 и 6, в сумме 7, бросаем еще раз."
	got = rollDiceMsg(1, 6, 7, 8, 1)
	assertMsg(t, got, want)

	want = "Выпало 5 и 3, в сумме 8, на это потребовалось 5 бросков."
	got = rollDiceMsg(5, 3, 8, 8, 5)
	assertMsg(t, got, want)
}

func TestRollDise(t *testing.T) {
	rollDice(8)
}
