package _7_3

import "errors"

const (
	MIN_NUM = 1
	MAX_NUM = 100
	MAX_TRY = 6
)

var val, cnt int

func guess(num int) (int, error) {
	cnt++
	if cnt > MAX_TRY {
		return 0, errors.New("too many attempts")
	}
	if num < val {
		return -1, nil
	}
	if num == val {
		return 0, nil
	}
	return 1, nil
}

func play() int {
	start := MIN_NUM - 1
	end := MAX_NUM
	for i := 0; i < MAX_TRY; i++ {
		num := (end + start) / 2
		if res, _ := guess(num); res == 1 {
			end = num - 1
		} else if res == -1 {
			start = num + 1
		} else {
			return num
		}
	}
	return (end + start) / 2
}
