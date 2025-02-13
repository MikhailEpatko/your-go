package _5_13

import (
	"errors"
	"fmt"
)

func fetchUserInfo(id string) (sum int, err error) {
	if id == "unknownId" {
		err = errors.New("not found")
	} else {
		sum = 12345
	}
	return
}

func UserProfile(id string) (msg string, err error) {
	if sum, e := fetchUserInfo(id); e != nil {
		err = fmt.Errorf("fetch error: %w", e)
	} else {
		msg = fmt.Sprintf("Пользователь с id %s имеет на счету %.2f руб.", id, float64(sum)/100)
	}
	return
}
