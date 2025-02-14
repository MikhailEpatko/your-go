package _7_1

import (
	"fmt"
	"math/rand/v2"
)

const (
	ROL_MSG    = "Выпало %d и %d, в сумме %d, бросаем еще раз."
	FINISH_MSG = "Выпало %d и %d, в сумме %d, на это потребовал%s %d брос%s."
	MIN        = 1
	MAX        = 6
)

func rollDice(goal int) {
	var cnt, sum int
	for sum != goal {
		cnt++
		res1 := rand.IntN(MAX) + MIN
		res2 := rand.IntN(MAX) + MIN
		sum = res1 + res2
		fmt.Println(rollDiceMsg(res1, res2, sum, goal, cnt))
	}
}

func rollDiceMsg(res1, res2, sum, goal, cnt int) string {
	if sum == goal {
		return fmt.Sprintf(FINISH_MSG, res1, res2, sum, suffix1(cnt), cnt, suffix2(cnt))
	} else {
		return fmt.Sprintf(ROL_MSG, res1, res2, sum)
	}
}

func suffix1(cnt int) string {
	if cnt%100 != 11 && cnt%10 == 1 {
		return "ся"
	}
	return "ось"
}

func suffix2(cnt int) string {
	if cnt < 5 {
		switch cnt {
		case 1:
			return "ок"
		case 2, 3, 4:
			return "ка"
		}
	}
	if cnt > 20 {
		switch cnt % 10 {
		case 1:
			return "ок"
		case 2, 3, 4:
			return "ка"
		}
	}
	return "ков"
}
