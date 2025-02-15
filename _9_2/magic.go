package _9_2

import (
	"fmt"
	"strconv"
	"strings"
)

func printMagic(sl []int) {
	fmt.Println(
		magicString(
			magicSlice(sl),
		),
	)
}

func magicString(numSlice []int) string {
	strSlice := make([]string, len(numSlice))
	for i, num := range numSlice {
		strSlice[i] = strconv.Itoa(num)
	}
	return "[" + strings.Join(strSlice, ", ") + "]"
}

func magicSlice(sl []int) []int {
	if len(sl) == 0 {
		return sl
	}
	res := make([]int, len(sl))
	for i := range sl {
		res[i] = mult(sl, i)
	}
	return res
}

func mult(sl []int, excludeIndex int) int {
	res := 1
	for i, n := range sl {
		if i == excludeIndex {
			continue
		}
		res *= n
	}
	return res
}
