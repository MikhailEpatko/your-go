package _9_4

func filterEven(v ...int) []int {
	res := []int{}
	for _, num := range v {
		if num%2 == 0 {
			res = append(res, num)
		}
	}
	return res
}
