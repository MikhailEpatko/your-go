package _5_18

func SumOfDigits(x int) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		x = -x
	}
	return (x % 10) + SumOfDigits(x/10)
}
