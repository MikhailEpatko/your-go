package _9_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinMax(t *testing.T) {
	a := assert.New(t)

	mi := []int{1, 2, 3}
	ma := []int{2, 3, 4}

	minWant, maxWant := mi, ma
	minGot, maxGot := minMax(ma, mi)
	a.Equal(minWant, minGot)
	a.Equal(maxWant, maxGot)
}

func TestIntersectSlices(t *testing.T) {
	a := assert.New(t)

	t.Run("when slices intersect should return all intersections", func(t *testing.T) {
		want := []int{3, 4, 5}
		got, err := intersectSlices(
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 6, 7},
		)
		a.Nil(err)
		a.Equal(want, got)
	})

	t.Run("when slices don't intersect should return empty slice", func(t *testing.T) {
		want := []int{}
		got, err := intersectSlices(
			[]int{1, 3, 5, 7},
			[]int{2, 4, 6, 8},
		)
		a.Nil(err)
		a.Equal(want, got)
	})

	t.Run("when slices are empty should return empty slice", func(t *testing.T) {
		want := []int{}
		got, err := intersectSlices(
			[]int{},
			[]int{},
		)
		a.Nil(err)
		a.Equal(want, got)
	})

	t.Run("when slice 2 is empty should return empty slice", func(t *testing.T) {
		want := []int{}
		got, err := intersectSlices(
			[]int{1, 2, 3},
			[]int{},
		)
		a.Nil(err)
		a.Equal(want, got)
	})

	t.Run("when slice 1 is empty should return empty slice", func(t *testing.T) {
		want := []int{}
		got, err := intersectSlices(
			[]int{},
			[]int{4, 5, 6},
		)
		a.Nil(err)
		a.Equal(want, got)
	})

	t.Run("when slice is nil should return error", func(t *testing.T) {
		var gotErr error
		wantErr := "slices cannot be nil"
		want := []int{}
		got, gotErr := intersectSlices(
			nil,
			[]int{2, 4, 6, 8},
		)
		a.NotNil(gotErr)
		a.Equal(wantErr, gotErr.Error())
		a.Equal(want, got)
	})
}
