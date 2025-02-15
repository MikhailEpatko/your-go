package _7_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTableStr(t *testing.T) {
	a := assert.New(t)

	want := "1 x 1 = 1\t1 x 2 = 2\t1 x 3 = 3\n2 x 1 = 2\t2 x 2 = 4\t2 x 3 = 6\n3 x 1 = 3\t3 x 2 = 6\t3 x 3 = 9\n"
	got := TableStr(3)
	a.Equal(want, got)
}
