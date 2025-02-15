package _9_5

import "fmt"

func Max(sl []int) (int, error) {
	if len(sl) == 0 {
		return 0, fmt.Errorf("slice is nil or empty")
	}
	m := sl[0]
	for _, n := range sl {
		m = max(m, n)
	}
	return m, nil
}
