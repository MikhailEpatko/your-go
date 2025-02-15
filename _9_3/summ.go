package _9_3

func SumSlices(sl1, sl2 []int) []int {
	var length int
	if len(sl1) < len(sl2) {
		length = len(sl1)
	} else {
		length = len(sl2)
	}
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = sl1[i] + sl2[i]
	}
	return result
}
