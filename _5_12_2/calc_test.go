package _5_12_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcResult(t *testing.T) {
	a := assert.New(t)

	t.Run("check add", func(t *testing.T) {
		got, err := Calculate(1.0, 2.0, "add")
		want := 3.0
		a.Nil(err)
		a.Equal(want, got)
	})
	t.Run("check subtract", func(t *testing.T) {
		got, err := Calculate(1.0, 2.0, "subtract")
		want := -1.0
		a.Nil(err)
		a.Equal(want, got)
	})
	t.Run("check multiply", func(t *testing.T) {
		got, err := Calculate(1.0, -2.0, "multiply")
		want := -2.0
		a.Nil(err)
		a.Equal(want, got)
	})
	t.Run("check divide", func(t *testing.T) {
		got, err := Calculate(3.0, -2.0, "divide")
		want := -1.5
		a.Nil(err)
		a.Equal(want, got)
	})
}

func TestCalcError(t *testing.T) {
	a := assert.New(t)

	t.Run("check zero error", func(t *testing.T) {
		res, err := Calculate(1.0, 0.0, "divide")
		want := "division by zero"
		a.Equal(0.0, res)
		a.NotNil(err)
		a.Equal(want, err.Error())
	})
	t.Run("check operation error", func(t *testing.T) {
		res, err := Calculate(1.0, 0.0, "qwerty")
		want := "unknown operation"
		a.Equal(0.0, res)
		a.NotNil(err)
		a.Equal(want, err.Error())
	})
}
