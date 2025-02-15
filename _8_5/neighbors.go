package _8_5

const (
	ARRAAY_LENGTH = 10
)

func SumNeighbors(arr [ARRAAY_LENGTH]int) [ARRAAY_LENGTH]int {
	res := [len(arr)]int{}
	for i := 0; i < len(arr); i++ {
		switch i {
		case 0:
			res[i] = arr[i+1]
		case len(arr) - 1:
			res[i] = arr[i-1]
		default:
			res[i] = arr[i-1] + arr[i+1]
		}
	}
	return res
}
