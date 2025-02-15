package _8_3

import "fmt"

const (
	SUCCESS_MSG = "Это палиндром!"
	FAIL_MSG    = "Не палиндром!"
)

func isPalindrome(nums [10]int) {
	fmt.Println(isPalindromeStr(nums))
}

func isPalindromeStr(nums [10]int) string {
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i] != nums[j] {
			return FAIL_MSG
		}
		i++
		j--
	}
	return SUCCESS_MSG
}
