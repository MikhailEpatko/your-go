package _7_3

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlay(t *testing.T) {
	a := assert.New(t)

	for i := 1; i <= 100; i++ {
		s := fmt.Sprintf("i = %d", i)
		t.Run(s, func(t *testing.T) {
			cnt = 0
			val = i
			got := play()
			//assertNum(t, val, got, cnt)
			a.Equal(val, got)
			a.True(cnt <= 6)
		})
	}
}
