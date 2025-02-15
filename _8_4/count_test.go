package _8_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountVowelsInArrayStr(t *testing.T) {
	a := assert.New(t)

	want := "3 2 2"
	got := countVowelsInArrayStr(&[STR_COUNT]string{"яблоко", "банан", "вишня"})
	a.Equal(want, got)
}
