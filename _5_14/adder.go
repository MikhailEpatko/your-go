package _5_14

func Adder(n int) func(int) int {
	sum := n
	return func(x int) int {
		sum += x
		return sum
	}
}
