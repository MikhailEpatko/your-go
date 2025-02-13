package _5_12_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcResult(t *testing.T) {
	a := assert.New(t)
	assertCalcResult := func(t testing.TB, got, want float64, err error) {
		t.Helper()
		a.Equal(want, got)
		a.Nil(err)
	}

	t.Run("check add", func(t *testing.T) {
		got, err := Calculate(1.0, 2.0, "add")
		want := 3.0
		assertCalcResult(t, got, want, err)
	})
	t.Run("check subtract", func(t *testing.T) {
		got, err := Calculate(1.0, 2.0, "subtract")
		want := -1.0
		assertCalcResult(t, got, want, err)
	})
	t.Run("check multiply", func(t *testing.T) {
		got, err := Calculate(1.0, -2.0, "multiply")
		want := -2.0
		assertCalcResult(t, got, want, err)
	})
	t.Run("check divide", func(t *testing.T) {
		got, err := Calculate(3.0, -2.0, "divide")
		want := -1.5
		assertCalcResult(t, got, want, err)
	})
}

func TestCalcError(t *testing.T) {
	a := assert.New(t)
	assertCalcResult := func(t testing.TB, res float64, want string, got error) {
		t.Helper()
		a.Equal(0.0, res)
		a.NotNil(got)
		a.Equal(want, got.Error())
	}

	t.Run("check zero error", func(t *testing.T) {
		res, got := Calculate(1.0, 0.0, "divide")
		want := "division by zero"
		assertCalcResult(t, res, want, got)
	})
	t.Run("check operation error", func(t *testing.T) {
		res, got := Calculate(1.0, 0.0, "qwerty")
		want := "unknown operation"
		assertCalcResult(t, res, want, got)
	})
}
