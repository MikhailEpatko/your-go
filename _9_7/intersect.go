package _9_7

import "fmt"

func intersectSlices(sl1 []int, sl2 []int) ([]int, error) {
	if sl1 == nil || sl2 == nil {
		return []int{}, fmt.Errorf("slices cannot be nil")
	}
	if len(sl1) == 0 || len(sl2) == 0 {
		return []int{}, nil
	}
	mini, maxi := minMax(sl1, sl2)
	res := make([]int, 0)
	for _, n := range maxi {
		if mini[len(mini)-1] < n {
			break
		}
		for _, m := range mini {
			if m == n {
				res = append(res, m)
				break
			}
		}
	}
	return res, nil
}

func minMax(sl1 []int, sl2 []int) (min []int, max []int) {
	if sl1[0] < sl2[0] {
		return sl1, sl2
	} else {
		return sl2, sl1
	}
}
