package _5_22

import (
	"errors"
	"fmt"
)

const (
	MOVE        = "Пират переместился на плиту %d"
	DONE        = "\nПират преодолел все ловушки"
	HURT        = "\nПират ранен"
	KILLED      = "\nПират убит"
	GAME_OVER   = "game over"
	EMPTY_MSG   = ""
	MAX_STEPS   = 3
	FATAL_HURTS = 2
)

var stepCnt, hurtCnt int

func MovePirate(isTrap bool) {
	if msg, err := movePirateMsg(isTrap); err == nil {
		fmt.Println(msg)
	}
}

func movePirateMsg(isTrap bool) (string, error) {
	stepCnt++
	if isTrap {
		hurtCnt++
	}
	if stepCnt > MAX_STEPS || hurtCnt > FATAL_HURTS {
		return EMPTY_MSG, errors.New(GAME_OVER)
	}
	msg := fmt.Sprintf(MOVE, stepCnt)
	if hurtCnt == FATAL_HURTS {
		msg += KILLED
	} else if isTrap {
		msg += HURT
	} else if stepCnt == MAX_STEPS {
		msg += DONE
	}
	return msg, nil
}
